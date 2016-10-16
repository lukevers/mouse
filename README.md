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
    pattern = "scripts/javascript/*.js"

    # Run plugins if these events happen
    events = [ "PRIVMSG" ]

    # Reload all JavaScript plugins or not on every event type listed above
    reload = true
```
