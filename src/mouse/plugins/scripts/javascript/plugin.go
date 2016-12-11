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

// A Plugin represents a JavaScript Plugin for a Mouse bot.
type Plugin struct {
	Config *Config
	Mouse  *mouse.Mouse

	event   *otto.Object
	storage *otto.Object

	files   []string
	vm      *otto.Otto
	mutex   *sync.Mutex
	watcher *fsnotify.Watcher
}

// NewPlugin creates a new JavaScript Plugin for a mouse bot.
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

	for _, folder := range plugin.Config.Folders {
		if err = plugin.watcher.Add(folder); err != nil {
			panic(err)
		}
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

	// Update global variables
	plugin.event.Set("command", event.Command)
	plugin.event.Set("channel", event.Channel)
	plugin.event.Set("message", event.Message)
	plugin.event.Set("host", event.Host)
	plugin.event.Set("nick", event.Nick)
	plugin.event.Set("user", event.User)

	plugin.storage.Set("event", plugin.event)

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
	plugin.mutex.Lock()
	defer plugin.mutex.Unlock()

	var files []string
	for _, folder := range plugin.Config.Folders {
		f, _ := filepath.Glob(fmt.Sprintf(
			"%s%s",
			folder,
			plugin.Config.Pattern,
		))

		files = append(files, f...)
	}

	plugin.files = files
	return
}

func (plugin *Plugin) register() {
	// Register event object
	plugin.event, _ = plugin.vm.Object("event = {}")
	plugin.vm.Set("event", plugin.event)

	// Register storage object and functions
	plugin.storage, _ = plugin.vm.Object("storage = {}")
	plugin.vm.Set("storage", plugin.storage)
	plugin.storage.Set("get", plugin.get)
	plugin.storage.Set("put", plugin.put)
	plugin.storage.Set("delete", plugin.delete)

	// Register global functions
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
