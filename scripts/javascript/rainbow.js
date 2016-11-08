function rainbow(str) {                              
    var count = 3;
    var short = false;

    if (str.length < 6) {
        short = true;
    }

    return str.replace(/./g, function(char) {
        if (char === ' ') {
            return char;
        }

        if (short) {
            count += 2;
        } else {
            if (++count > 12) {
                count = 3;
            }
        }

        return '\x03' + count + ',99' + char;
    });
}

if (event.command == 'PRIVMSG') {
    if (event.message.indexOf('| rainbow') > -1) {
        var message = rainbow(event.message.replace('| rainbow', '').replace(/\s+/g, ' '));
        say(event.channel, message);
    } else if (event.message.indexOf('@rainbow') > -1) {
        var message = rainbow(event.message.replace('@rainbow', '').replace(/\s+/g, ' '));
        say(event.channel, message.trim());
    }
}
