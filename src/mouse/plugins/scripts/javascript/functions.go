package javascript

import (
	"github.com/robertkrimen/otto"
)

func (plugin *Plugin) join(call otto.FunctionCall) otto.Value {
	return otto.Value{}
}

func (plugin *Plugin) part(call otto.FunctionCall) otto.Value {
	return otto.Value{}
}

func (plugin *Plugin) cycle(call otto.FunctionCall) otto.Value {
	return otto.Value{}
}

func (plugin *Plugin) say(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	message, _ := call.Argument(1).ToString()

	plugin.Mouse.Say(channel, message)

	return otto.Value{}
}

func (plugin *Plugin) kick(call otto.FunctionCall) otto.Value {
	return otto.Value{}
}

func (plugin *Plugin) ban(call otto.FunctionCall) otto.Value {
	return otto.Value{}
}
