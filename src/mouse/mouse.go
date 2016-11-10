package mouse

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/sorcix/irc.v1"
	"net"
	"sync"
	"time"
)

type Mouse struct {
	Config *Config

	conn   net.Conn
	reader *irc.Decoder
	writer *irc.Encoder
	mutex  *sync.Mutex

	data     chan *irc.Message
	handlers []func(*Event)
}

func New(config Config) *Mouse {
	return &Mouse{
		Config: &config,
		data:   make(chan *irc.Message, 10),
		mutex:  &sync.Mutex{},
	}
}

func (mouse *Mouse) Connect() error {
	var err error
	server := fmt.Sprintf("%s:%d", mouse.Config.Host, mouse.Config.Port)

	if mouse.Config.TLS {
		mouse.conn, err = tls.Dial("tcp", server, mouse.Config.TLSConfig)
	} else {
		mouse.conn, err = net.Dial("tcp", server)
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

	go mouse.loop()

	return err
}

func (mouse *Mouse) loop() {
	go mouse.handle()

	for {
		mouse.conn.SetDeadline(time.Now().Add(300 * time.Second))
		message, err := mouse.reader.Decode()
		if err != nil {
			break
		}

		mouse.data <- message
	}

	if mouse.Config.Reconnect {
		mouse.Connect()
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

func (mouse *Mouse) pong(message *irc.Message) {
	mouse.writer.Encode(&irc.Message{
		Command: irc.PONG,
		Params:  []string{message.Trailing},
	})
}

func (mouse *Mouse) Use(handler func(*Event)) {
	mouse.handlers = append(mouse.handlers, handler)
}

func (mouse *Mouse) Join(channel string) error {
	return mouse.writer.Encode(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{channel},
	})
}

func (mouse *Mouse) Part(channel string) error {
	return mouse.writer.Encode(&irc.Message{
		Command: irc.PART,
		Params:  []string{channel},
	})
}

func (mouse *Mouse) Say(channel, message string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.PRIVMSG,
		Params:   []string{channel},
		Trailing: message,
	})
}

func (mouse *Mouse) Op(channel, nick string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.MODE,
		Params:   []string{channel, "+o"},
		Trailing: nick,
	})
}

func (mouse *Mouse) Deop(channel, nick string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.MODE,
		Params:   []string{channel, "-o"},
		Trailing: nick,
	})
}

func (mouse *Mouse) Kick(channel, user, reason string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.KICK,
		Params:   []string{user},
		Trailing: reason,
	})
}

func (mouse *Mouse) Ban(channel, user, reason string) error {
	return mouse.writer.Encode(&irc.Message{
		Command:  irc.MODE,
		Params:   []string{channel, "+b", user},
		Trailing: reason,
	})
}

func (mouse *Mouse) Unban(channel, user string) error {
	return mouse.writer.Encode(&irc.Message{
		Command: irc.MODE,
		Params:  []string{channel, "-b", user},
	})
}
