if (event.command == 'PRIVMSG') {
    run();
}

function run() {
    if (event.message == '@ping') {
        say(event.channel, event.nick + ': pong');
    }
}
