function Roulette() {
    this.position = storage.get('position');
    this.bullet = storage.get('bullet');
    this.counter = storage.get('counter');

    this.empty = storage.get('empty');
    if (this.empty == '') {
        this.empty = 'true';
        storage.put('empty', 'true');
    }
}

Roulette.prototype.invoked = function invoked() {
    return event.message.indexOf('!r') == 0;
};

Roulette.prototype.command = function command() {
    var params = event.message.trim().split(' ');
    return params.length === 1 ? 'run' : params[1];
};

Roulette.prototype.reload = function reload() {
    if (this.empty == 'false') {
        this.message('I\'m already loaded.');
    } else {
        storage.put('bullet', Math.floor(Math.random() * 6) + 1);
        storage.put('position', 0);
        storage.put('empty', false);
        storage.put('counter', 0);
    }
};

Roulette.prototype.spin = function spin() {
    if (this.empty == 'true') {
        this.message('Weee. Nothing to spin though.');
    } else {
        this.bullet = Math.floor(Math.random() * 6) + 1;
        storage.put('bullet', this.bullet);
    }
};

Roulette.prototype.run = function run() {
    if (this.empty == 'true') {
        this.message('Nothing to shoot. Reload first.');
    } else {
        if (++this.position > 6) {
            this.position = 0;
        }

        if (this.position == this.bullet) {
            storage.put('empty', true);
            kick(event.channel, event.nick, 'BOOM');
        } else {
            // Increment position & counter
            storage.put('position', this.position);
            storage.put('counter', ++this.counter);
            this.message(this.counter);
        }
    }
};

Roulette.prototype.message = function message(message) {
    say(event.channel, event.nick + ': ' + message);
};

if (event.command == 'PRIVMSG') {
    var r = new Roulette();

    if (r.invoked()) {
        switch (r.command()) {
            case 'load':
            case 'reload':
                r.reload();
                break;
            case 'spin':
                r.spin();
                break;
            case 'run':
                r.run();
                break;
            default:
                r.message('!r [reload|spin|run]');
                break;
        }
    }
}
