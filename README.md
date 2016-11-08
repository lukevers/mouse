![Mouse](./vendor/mouse.jpg)

A scriptable, configuration powered IRC bot that can handle as many connnections as you want.

## Table of Contents

1. [Installing](#installing)
    1. [Building from source](#building-from-source)
    2. [Downloading a binary](#downloading-a-binary)
2. [Configuring](#configuring)
    1. [Configuring with TOML](#configuring-with-toml)
    2. [Configuring with JSON](#configuring-with-json)
    3. [Configuring with HCL](#configuring-with-hcl)
    4. [Configuring with YAML](#configuring-with-yaml)
    5. [Configuring with Java Properties](#configuring-with-java-properties)
3. [Extending Mouse with plugins](#extending-mouse-with-plugins)
    1. [Writing a plugin in JavaScript](#writing-a-plugin-in-javascript)
        1. [Global JavaScript functions](#global-javascript-functions)
            1. [Join](#join)
            2. [Part](#part)
            3. [Cycle](#cycle)
            4. [Say](#say)
            5. [Kick](#kick)
            6. [Ban](#ban)
        2. [Global JavaScript data](#global-javascript-data)
            1. [Event](#event)
                1. [Command](#command)
                2. [Channel](#channel)
                3. [Message](#message)
                4. [Host](#host)
                5. [Nick](#nick)
                6. [User](#user)
4. [Contributing](#contributing)
5. [License](#license)

## Installing

### Building from source

Mouse is built on [Go](https://golang.org/), and uses [GB](https://getgb.io/) to vendor code. You should always be using the most recent version of both.

```bash
gb build all
```

### Downloading a binary

This option is currently not available, but will be once Mouse hits version `1.0.0`, and for each release after that.

## Configuring

Mouse uses the [Viper](https://github.com/spf13/viper) library, which allows a variety of configuration types. Your configuration file can exist at any of the following locations:

```bash
/etc/mouse/
$HOME/.mouse/
./
```

Your configuration file must be named appropriately and should be one of the following:

```bash
mouse.toml
mouse.json
mouse.hcl
mouse.yaml
mouse.properties
```

Keep in mind that a configuration file named `mouse.toml` MUST be a TOML file, and the same goes for every other supported configuration file type.

### Configuring with TOML

You can choose to configure Mouse with [TOML](https://github.com/toml-lang/toml). Configure your servers like this:

```toml
[servers]

    [server.a]
        nick = "mouse"

        # ...

        [servers.a.plugins.javascript]
            enabled = true

            # ...

    [server.b]
        nick = "mouse"

        # ...
```

You can see a full example at [contrib/config-examples/config.toml](contrib/config-examples/config.toml).

### Configuring with JSON

You can choose to configure Mouse with [JSON](http://www.json.org/). Configure your servers like this:

```json
{
    "servers": {
        "a": {
            "nick": "mouse",
            "plugins": {
                "javascript": {
                    "enabled": true
                }
            }
        },
        "b": {
            "nick": "mouse"
        }
    }
}
```

You can see a full example at [contrib/config-examples/config.json](contrib/config-examples/config.json).

### Configuring with HCL

You can choose to configure Mouse with [HCL](https://github.com/hashicorp/hcl). Configure your servers like this:

```hcl
servers "a" {
    nick = "mouse"

    # ...

    plugins "javascript" {
        enabled = true

        # ...
    }
}

servers "b" {
    nick = "mouse"

    # ...
}
```

You can see a full example at [contrib/config-examples/config.hcl](contrib/config-examples/config.hcl).

### Configuring with YAML

You can choose to configure Mouse with [YAML](http://yaml.org/). Configure your servers like this:

```yaml
servers:
    a:
        nick: mouse
        # ...

        plugins:
            javascript:
                enabled: true
                # ...
```

You can see a full example at [contrib/config-examples/config.yaml](contrib/config-examples/config.yaml).

### Configuring with Java Properties

You can choose to configure Mouse with [Java Properties](http://docs.oracle.com/javase/tutorial/essential/environment/properties.html).

```properties
servers.a.nick = "mouse"
# ...
servers.a.plugins.javascript.enabled = true
# ...

servers.b.nick = "mouse"
# ...
```

You can see a full example at [contrib/config-examples/config.properties](contrib/config-examples/config.properties).

## Extending Mouse with plugins

TODO

### Writing a plugin in JavaScript

TODO

#### Global JavaScript functions

##### Join

TODO

```javascript
/**
 * @param string channel
 */
function join(channel)
```

##### Part

TODO

```javascript
/**
 * @param string channel
 */
function part(channel)
```

##### Cycle

TODO

```javascript
/**
 * @param string channel
 */
function cycle(channel)
```

##### Say

The `say` function allows your bot to send messages to any buffer that will allow it. When sending to a channel, the channel name should be prefixed with `#`, and when sending to a user, it should not be.

```javascript
/**
 * @param string buffer
 * @param string message
 */
function say(buffer, message)
```

##### Kick

TODO

```javascript
/**
 * @param string channel
 * @param string user
 */
function kick(channel, user)
```

##### Ban

TODO

```javascript
/**
 * @param string channel
 * @param string user
 */
function ban(channel, user)
```

#### Global JavaScript data

There are variables set at the global scope that can be used.

##### Event

On each event that the JavaScript plugins are listening for, a new event is populated that is passed in to an `event` object in the global scope.

###### Command

Command is a string that contains the command that triggered this event. The most frequently used command in Mouse plugins is probably [`"PRIVMSG"`](https://tools.ietf.org/html/rfc2812#section-3.3.1).

```javascript
event.command
```

###### Channel

Channel is a string that contains the channel where the event took place. Channel is not just for channels, but also for private messages. If it is a channel, it will be prefixed with `#`.

```javascript
event.channel
```

###### Message

Message is a string that contains the message of the event.

```javascript
event.message
```

###### Host

Host is a string that contains the host of the user that triggered this event.

```javascript
event.host
```

###### Nick

Nick is a string that contains the nick of the user that triggered this event.

```javascript
event.nick
```

###### User

User is a string that contains the user of the user that triggered this event.

```javascript
event.user
```

## Contributing

Want to contribute? Do it. Read the [contributing guidelines](CONTRIBUTING.md) first.

## License

Mouse is licensed under [MIT](LICENSE.md).
