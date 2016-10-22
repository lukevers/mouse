var messages = {
    "hello": [
        "hey {0}",
        "hi {0}",
        "hello {0}",
        "hai {0}",
        "{0}: hey",
        "{0}: hi",
        "{0}: hello",
        "{0}: hai",
    ],
    "thanks": [
        "thanks {0}",
        "thank you {0}",
        "{0}: thanks",
    ],
};

var responses = {
    "hello": [
        "hey " + irc.event.nick,
        "hi " + irc.event.nick,
        "hello " + irc.event.nick,
        "hai " + irc.event.nick,
    ],
    "thanks": [
        "you're welcome " + irc.event.nick,
        irc.event.nick + " you're welcome",
    ],
};

var name = "systemd";

String.prototype.format = function() {
    var formatted = this;
    for (var i = 0; i < arguments.length; i++) {
        var regexp = new RegExp('\\{'+i+'\\}', 'gi');
        formatted = formatted.replace(regexp, arguments[i]);
    }

    return formatted;
};

if (irc.event.command == 'PRIVMSG') {
    for (var index in messages) {
        if (messages.hasOwnProperty(index)) {
            for (var i = 0; i < messages[index].length; i++) {
                messages[index][i] = messages[index][i].format(name, '{0}');
                if (irc.event.message.indexOf(messages[index][i]) > -1) {
                    var r = [Math.floor(Math.random() * responses[index].length)];
                    irc.say(irc.event.channel, responses[index][r]);
                    break;
                }
            }
        }
    }
}
