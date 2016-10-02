package javascript

import (
	"github.com/robertkrimen/otto"
	"mouse"
	"path/filepath"
)

type Plugin struct {
	Config *Config
	Mouse  *mouse.Mouse

	files []string
	vm    *otto.Otto
	irc   *otto.Object
	event *otto.Object
}

func NewPlugin(mouse *mouse.Mouse, config *Config) func(*mouse.Event) {
	plugin := &Plugin{
		Config: config,
		Mouse:  mouse,
		vm:     otto.New(),
	}

	// Register javascript functions
	plugin.register()

	// Load initial scripts
	if err := plugin.load(); err != nil {
		panic(err)
	}

	return plugin.handler
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

	// Reload plugins each time if we have ContinuousLoad set to true.
	if plugin.Config.ContinuousLoad {
		if err := plugin.load(); err != nil {
			panic(err)
		}
	}

	// Setup the event in the global IRC object
	plugin.event.Set("command", event.Command)
	plugin.event.Set("channel", event.Channel)
	plugin.event.Set("message", event.Message)
	plugin.event.Set("host", event.Host)
	plugin.event.Set("nick", event.Nick)
	plugin.event.Set("user", event.User)

	for _, file := range plugin.files {
		script, err := plugin.vm.Compile(file, nil)
		if err != nil {
			panic(err)
		}

		plugin.vm.Run(script)
	}
}

func (plugin *Plugin) load() (err error) {
	plugin.files, err = filepath.Glob(plugin.Config.Pattern)
	return
}

func (plugin *Plugin) register() {
	// Create a global IRC object
	plugin.irc, _ = plugin.vm.Object("irc = {}")
	plugin.event, _ = plugin.vm.Object("event = {}")

	// Register functions to the IRC object
	plugin.irc.Set("say", plugin.say)

	// Register the global IRC object
	plugin.irc.Set("event", plugin.event)
	plugin.vm.Set("irc", plugin.irc)
}

func (plugin *Plugin) say(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	message, _ := call.Argument(1).ToString()

	plugin.Mouse.Say(channel, message)

	return otto.Value{}
}
