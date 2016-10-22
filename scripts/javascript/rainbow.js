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

        return '\x03' + count + char;      
    });
}

if (irc.event.command == 'PRIVMSG') {
	if (irc.event.message.indexOf('| rainbow') > -1) {
		var message = rainbow(irc.event.message.replace('| rainbow', '').replace(/\s+/g, ' '));
		irc.say(irc.event.channel, message);
	} else if (irc.event.message.indexOf('@rainbow') > -1) {
		var message = rainbow(irc.event.message.replace('@rainbow', '').replace(/\s+/g, ' '));
		irc.say(irc.event.channel, message.trim());
	}
}
