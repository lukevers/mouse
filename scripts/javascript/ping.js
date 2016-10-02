if (irc.event.command == 'PRIVMSG') {
    run();
}

function run() {
    if (irc.event.message == '@ping') {
        irc.say(irc.event.channel, irc.event.nick + ': pong');
    }
}
