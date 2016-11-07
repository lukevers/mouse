![Mouse](./vendor/mouse.jpg)

A scriptable, configuration powered IRC bot that can handle as many connnections as you want.

## Configuring Mouse

Your configuration file can either be written in [TOML](https://github.com/toml-lang/toml), [JSON](http://www.json.org/), [YAML](http://yaml.org/), [HCL](https://github.com/hashicorp/hcl), or [Java properties](http://docs.oracle.com/javase/tutorial/essential/environment/properties.html). Mouse will search for your configuration file at all of these locations with the name `config.{toml|json|yaml|hcl|properties}`:

```bash
/etc/mouse/
$HOME/.mouse/
./
```

You can see an example of each file type in the `examples` folder.

## Writing Plugins

At this time, Mouse only supports JavaScript plugins, but will soon also support Lua and possibly other plugins.

### JavaScript Plugins

Writing a JavaScript plugin is easy. There are some examples in `scripts/javascript/*.js` if you want to see some examples for yourself. Mouse adds a global `irc` object which contains both data from the most recent message and functions to interact with the IRC server.

Here's a pseducode example of what the global `irc` object looks like and what data/functions are available:

```javascript
{
    // Updates on every message
    event: {
        command: "PRIVMSG",
        channel: "#channel",
        message: "hey what's up",

        // Information about the sender
        host: "irc.lukevers.com",
        nick: "lukevers",
        user: "lukevers"
    }

    say: function(channel, message) {
        // Send a PRIVMSG message to a channel or user.
    },
}
```
