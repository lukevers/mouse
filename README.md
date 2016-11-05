# Mouse

A scriptable/configurable IRC bot that can handle as many connections as you want.

## Configuring the mice

The config file is written in [TOML](https://github.com/toml-lang/toml). You can see an example at [`config.example.toml`](./config.example.toml).

Each bot (mouse) lives in a separate section in a global `[servers]` key:

```toml
[servers]

    [servers.server_a]
        # ...

    [servers.server_b]
        # ...
```

Each bot has the following configuration options:

```toml
nick = "somenick"
user = "someuser"
name = "somename"

host = "irc.freenode.net"
port = 6667

# Currently does nothing because I haven't setup the config options for TLS
tls = false

# List of channels to join on connect
channels = [ "#channel", "#other-channel" ]

# Logs all events to STDOUT
debug = true
```

And each bot has the following configuration options for plugins:

```toml
[servers.server_a.plugins.javascript]
    # Enable JavaScript plugins or not
    enabled = true

    # Pattern to where the scripts are
    folder = "scripts/javascript/"
    pattern = "*.js"

    # Run plugins if these events happen
    events = [ "PRIVMSG" ]

    # Reload all JavaScript plugins or not on every event type listed above
    reload = true
```

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
