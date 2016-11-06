server "fc00" {
    nick = "mouse"
    user = "mouse"
    name = "mouse"

    host = "irc.fc00.io"
    port = 6667
    tls = false
    reconnect = true
    debug = true

    channels = [ "#lukevers", "#mouse", "#mice" ]

    plugins "javascript" {
        enabled = true
        folder = "scripts/javascript/"
        pattern = "*.js"
        events = [ "PRIVMSG" ]
    }
}
