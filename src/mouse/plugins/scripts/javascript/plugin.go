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

	// Update the gloval event
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
	// Register data
	plugin.event, _ = plugin.vm.Object("event = {}")
	plugin.vm.Set("event", plugin.event)

	// Register functions
	plugin.vm.Set("join", plugin.join)
	plugin.vm.Set("part", plugin.part)
	plugin.vm.Set("cycle", plugin.cycle)
	plugin.vm.Set("say", plugin.say)
	plugin.vm.Set("kick", plugin.kick)
	plugin.vm.Set("ban", plugin.ban)
	plugin.vm.Set("unban", plugin.unban)
	plugin.vm.Set("op", plugin.op)
	plugin.vm.Set("deop", plugin.deop)
}
