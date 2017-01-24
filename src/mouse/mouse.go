package mouse

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/sorcix/irc.v1"
	"net"
	"storage"
	"sync"
	"time"
)

// A Mouse represents an IRC bot
type Mouse struct {
	Config  *Config
	Storage *storage.Store

	alive  int
	conn   net.Conn
	reader *irc.Decoder
	writer *irc.Encoder
	mutex  *sync.Mutex

	data     chan *irc.Message
	handlers []func(*Event)
}

// New creates a new IRC bot. If there are any problems creating the bot, or
// connecting to storage, this returns an error.
func New(config Config) (*Mouse, error) {
	mouse := Mouse{
		Config: &config,
		alive:  ConnectionWaiting,
		data:   make(chan *irc.Message, 10),
		mutex:  &sync.Mutex{},
	}

	var err error
	mouse.Storage, err = storage.New(config.StorageDriver, config.Storage.DSN)

	return &mouse, err
}

// Connect creates a new connection to the IRC server defined in the Config. If
// there are any problems connecting, this will return an error.
func (mouse *Mouse) Connect() error {
	var err error
	server := fmt.Sprintf("%s:%d", mouse.Config.Host, mouse.Config.Port)

	if mouse.Config.TLS {
		mouse.conn, err = tls.Dial("tcp", server, mouse.Config.TLSConfig)
	} else {
		mouse.conn, err = net.DialTimeout("tcp", server, 10*time.Second)
	}

	if err != nil {
		return err
	}

	// Create readers and writers from the connection
	mouse.reader = irc.NewDecoder(mouse.conn)
	mouse.writer = irc.NewEncoder(mouse.conn)

	// Write initial connect messages
	if mouse.Config.Pass != "" {
		if err = mouse.writer.Encode(&irc.Message{
			Command: irc.PASS,
			Params:  []string{mouse.Config.Pass},
		}); err != nil {
			return err
		}
	}

	if err = mouse.writer.Encode(&irc.Message{
		Command: irc.NICK,
		Params:  []string{mouse.Config.Nick},
	}); err != nil {
		return err
	}

	if err = mouse.writer.Encode(&irc.Message{
		Command: irc.USER,
		Params:  []string{mouse.Config.User, "0", "*", mouse.Config.Name},
	}); err != nil {
		return err
	}

	mouse.alive = ConnectionAlive
	go mouse.loop()

	// Join channels
	for _, channel := range mouse.Config.Channels {
		mouse.Join(channel)
	}

	return nil
}

func (mouse *Mouse) loop() {
	go mouse.handle()
	go mouse.checkConnection()

	for {
		if mouse.alive == ConnectionDead {
			break
		}

		mouse.conn.SetDeadline(time.Now().Add(30 * time.Second))
		message, err := mouse.reader.Decode()
		if err != nil {
			break
		}

		mouse.data <- message
	}

	if mouse.Config.Reconnect {
		for {
			if err := mouse.Connect(); err != nil {
				time.Sleep(10 * time.Second)
			} else {
				break
			}
		}
	}
}

func (mouse *Mouse) checkConnection() {
	for {
		// Set the connection to waiting
		mouse.alive = ConnectionWaiting

		// Ping the server to check the connection
		mouse.ping(mouse.Config.Host)

		// Sleep for however long the configuration option for Ping is set for
		time.Sleep(mouse.Config.Ping * time.Second)

		if mouse.alive == ConnectionWaiting {
			mouse.alive = ConnectionDead
			break
		}
	}
}

func (mouse *Mouse) handle() {
	for {
		// Receive the data
		message := <-mouse.data

		// If the command is a PING from the server, we need to respond ASAP.
		// We're responding in a goroutine so we can also pass the event on to
		// all of the event handlers that might rely on a PING event.
		if message.Command == irc.PING {
			go mouse.pong(message)
		}

		// If we recieved a PONG message, our connection is alive.
		if message.Command == irc.PONG {
			mouse.alive = ConnectionAlive
		}

		// Sometimes there is no channel, and when that happens it means it's a
		// message from the server. We set it to our current nick so we know.
		channel := mouse.Config.Nick
		if len(message.Params) > 0 {
			channel = message.Params[0]
		}

		// During some events (like "PING") Prefix is not set, which causes a
		// problem when we try to access it. We set it here so we don't panic.
		if message.Prefix == nil {
			message.Prefix = &irc.Prefix{}
		}

		event := &Event{
			Channel: channel,
			Command: message.Command,
			Host:    message.Prefix.Host,
			Message: message.Trailing,
			Nick:    message.Prefix.Name,
			User:    message.Prefix.User,
		}

		go func(event *Event) {
			mouse.mutex.Lock()
			defer mouse.mutex.Unlock()

			for _, handler := range mouse.handlers {
				handler(event)
			}

		}(event)
	}
}

func (mouse *Mouse) ping(server string) {
	mouse.writer.Encode(&irc.Message{
		Command: irc.PING,
		Params:  []string{server},
	})
}

func (mouse *Mouse) pong(message *irc.Message) {
	mouse.writer.Encode(&irc.Message{
		Command: irc.PONG,
		Params:  []string{message.Trailing},
	})
}

// Use adds a handler function to the stack of handlers for IRC events.
func (mouse *Mouse) Use(handler func(*Event)) {
	mouse.handlers = append(mouse.handlers, handler)
}

// Join allows the bot to join a channel. If there is a password for the
// channel the bot is trying to join, append it with a space after the channel
// name.
func (mouse *Mouse) Join(channel string) error {
	return mouse.writer.Encode(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{channel},
	})
}

// Part allows the bot to part a channel.
func (mouse *Mouse) Part(channel string) error {
	return mouse.writer.Encode(&irc.Message{
		Command: irc.PART,
		Params:  []string{channel},
	})
}

// Say allows the bot to send a PRIVMSG to a channel.
func (mouse *Mouse) Say(channel, message string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   []string{channel},
		Trailing: message,
	})
}

// Op allows the bot to change the mode to +o of a user given in a specific
// channel.
func (mouse *Mouse) Op(channel, nick string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.MODE,
		Params:   []string{channel, "+o"},
		Trailing: nick,
	})
}

// Deop allows the bot to change the mode to -o of a user given in a specific
// channel.
func (mouse *Mouse) Deop(channel, nick string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.MODE,
		Params:   []string{channel, "-o"},
		Trailing: nick,
	})
}

// Kick allows the bot to kick a user out of a specific channel. While you must
// pass a parameter as a reason, if you don't want to send an actual reason,
// you can pass an empty string.
func (mouse *Mouse) Kick(channel, user, reason string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.KICK,
		Params:   []string{channel, user},
		Trailing: reason,
	})
}

// Ban allows the bot to ban a user out of a specific channel. While you must
// pass a parameter as a reason, if you don't want to send an actual reason,
// you can pass an empty string.
func (mouse *Mouse) Ban(channel, user, reason string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.MODE,
		Params:   []string{channel, "+b", user},
		Trailing: reason,
	})
}

// Unban allows the bot to unban a user in a specific channel.
func (mouse *Mouse) Unban(channel, user string) error {
	return mouse.writer.Encode(&irc.Message{
		Command: irc.MODE,
		Params:  []string{channel, "-b", user},
	})
}
