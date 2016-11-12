servers "fc00" {
    nick = "mouse"
    user = "mouse"
    name = "mouse"

    host = "irc.fc00.io"
    port = 6667
    tls = false
    reconnect = true
    debug = true
    storage = "sqlite3"

    channels = [ "#lukevers", "#mouse", "#mice" ]

    plugins "javascript" {
        enabled = true
        folders = [ "contrib/scripts/javascript/" ]
        pattern = "*.js"
        events = [ "PRIVMSG" ]
    }

    store "sqlite3" {
        dsn = "/path/to/dbname.db"
        table = "fc00"
    }

    store "mysql" {
        dsn = "user:password@/dbname?charset=utf8&parseTime=True"
        table = "fc00"
    }

    store "postgres" {
        dsn = "host=myhost user=mouse dbname=mouse sslmode=disable password=mypassword"
        table = "fc00"
    }
}
