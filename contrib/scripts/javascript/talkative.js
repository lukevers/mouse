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
    "work-better": [
        "{0}: work better",
        "work better, {0}",
        "work better {0}",
    ],
};

var responses = {
    "hello": [
        "hey " + event.nick,
        "hi " + event.nick,
        "hello " + event.nick,
        "hai " + event.nick,
    ],
    "thanks": [
        "you're welcome " + event.nick,
        event.nick + " you're welcome",
    ],
    "work-better": [
        "ok",
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

if (event.command == 'PRIVMSG') {
    for (var index in messages) {
        if (messages.hasOwnProperty(index)) {
            for (var i = 0; i < messages[index].length; i++) {
                messages[index][i] = messages[index][i].format(name, '{0}');
                if (event.message.indexOf(messages[index][i]) > -1) {
                    var r = [Math.floor(Math.random() * responses[index].length)];
                    say(event.channel, responses[index][r]);
                    break;
                }
            }
        }
    }
}
