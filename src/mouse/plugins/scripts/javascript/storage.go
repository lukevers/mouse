package javascript

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"strings"
)

func (plugin *Plugin) getTable(call otto.FunctionCall) string {
	event, _ := call.This.Object().Get("event")
	ch, _ := event.Object().Get("channel")
	channel, _ := ch.ToString()

	parts := strings.Split(call.CallerLocation(), "/")
	part := parts[len(parts)-1]

	return fmt.Sprintf(
		"%s_%s_%s",
		plugin.Config.Name,
		strings.Trim(channel, "#"),
		strings.ToLower(strings.Split(part, ".js")[0]),
	)
}

func (plugin *Plugin) get(call otto.FunctionCall) otto.Value {
	key, _ := call.Argument(0).ToString()

	val := plugin.Config.Storage.Get(plugin.getTable(call), key)
	value, _ := otto.ToValue(val)

	return value
}

func (plugin *Plugin) put(call otto.FunctionCall) otto.Value {
	key, _ := call.Argument(0).ToString()
	val, _ := call.Argument(1).ToString()

	plugin.Config.Storage.Put(plugin.getTable(call), key, val)
	return otto.Value{}
}

func (plugin *Plugin) delete(call otto.FunctionCall) otto.Value {
	key, _ := call.Argument(0).ToString()

	plugin.Config.Storage.Delete(plugin.getTable(call), key)
	return otto.Value{}
}
