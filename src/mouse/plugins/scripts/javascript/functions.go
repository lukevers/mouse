package javascript

import (
	"github.com/robertkrimen/otto"
)

func (plugin *Plugin) join(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()

	plugin.Mouse.Join(channel)

	return otto.Value{}
}

func (plugin *Plugin) part(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()

	plugin.Mouse.Part(channel)

	return otto.Value{}
}

func (plugin *Plugin) cycle(call otto.FunctionCall) otto.Value {
	plugin.part(call)
	plugin.join(call)
	return otto.Value{}
}

func (plugin *Plugin) say(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	message, _ := call.Argument(1).ToString()

	plugin.Mouse.Say(channel, message)

	return otto.Value{}
}

func (plugin *Plugin) kick(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	user, _ := call.Argument(1).ToString()
	reason, _ := call.Argument(2).ToString()

	plugin.Mouse.Kick(channel, user, reason)

	return otto.Value{}
}

func (plugin *Plugin) ban(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	user, _ := call.Argument(1).ToString()
	reason, _ := call.Argument(2).ToString()

	plugin.Mouse.Ban(channel, user, reason)

	return otto.Value{}
}

func (plugin *Plugin) op(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	user, _ := call.Argument(1).ToString()

	plugin.Mouse.Op(channel, user)

	return otto.Value{}
}

func (plugin *Plugin) deop(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	user, _ := call.Argument(1).ToString()

	plugin.Mouse.Deop(channel, user)

	return otto.Value{}
}
