package javascript

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/robertkrimen/otto"
	"logger/stderr"
	"mouse"
	"path/filepath"
	"sync"
)

type Plugin struct {
	Config *Config
	Mouse  *mouse.Mouse

	files   []string
	vm      *otto.Otto
	irc     *otto.Object
	event   *otto.Object
	mutex   *sync.Mutex
	watcher *fsnotify.Watcher
}

func NewPlugin(mouse *mouse.Mouse, config *Config) func(*mouse.Event) {
	plugin := &Plugin{
		Config: config,
		Mouse:  mouse,
		mutex:  &sync.Mutex{},
		vm:     otto.New(),
	}

	// Register javascript functions
	plugin.register()
	go plugin.watchFiles()

	// Load initial scripts
	if err := plugin.load(); err != nil {
		stderr.Printf("Could not load javascript plugins:", err)
	}

	return plugin.handler
}

func (plugin *Plugin) watchFiles() {
	var err error
	plugin.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer plugin.watcher.Close()

	if err = plugin.watcher.Add(plugin.Config.Folder); err != nil {
		panic(err)
	}

	for {
		select {
		case _ = <-plugin.watcher.Events:
			if err = plugin.load(); err != nil {
				stderr.Printf("Could not reload javascript plugins:", err)
			}
		case err := <-plugin.watcher.Errors:
			stderr.Printf("Error occurred while watching javascript plugins:", err)
		}
	}
}

func (plugin *Plugin) handler(event *mouse.Event) {
	skip := true
	for _, eventType := range plugin.Config.EventTypes {
		if event.Command == eventType {
			skip = false
			break
		}
	}

	if skip {
		return
	}

	// Setup the event in the global IRC object
	plugin.event.Set("command", event.Command)
	plugin.event.Set("channel", event.Channel)
	plugin.event.Set("message", event.Message)
	plugin.event.Set("host", event.Host)
	plugin.event.Set("nick", event.Nick)
	plugin.event.Set("user", event.User)

	for _, file := range plugin.files {
		plugin.mutex.Lock()
		script, err := plugin.vm.Compile(file, nil)
		if err != nil {
			stderr.Printf("Error occurred while compiling javascript plugin:", err)
		}

		plugin.vm.Run(script)
		plugin.mutex.Unlock()
	}
}

func (plugin *Plugin) load() (err error) {
	plugin.files, err = filepath.Glob(fmt.Sprintf(
		"%s%s",
		plugin.Config.Folder,
		plugin.Config.Pattern,
	))

	return
}

func (plugin *Plugin) register() {
	// Create a global IRC object
	plugin.irc, _ = plugin.vm.Object("irc = {}")
	plugin.vm.Set("irc", plugin.irc)

	// Create irc.event which is populated on each run
	plugin.event, _ = plugin.vm.Object("event = {}")
	plugin.irc.Set("event", plugin.event)

	// Register functions to the IRC object
	plugin.irc.Set("say", plugin.say)
}

func (plugin *Plugin) say(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	message, _ := call.Argument(1).ToString()

	plugin.Mouse.Say(channel, message)

	return otto.Value{}
}
