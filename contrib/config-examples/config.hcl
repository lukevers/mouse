servers "fc00" {
    nick = "mouse"
    user = "mouse"
    name = "mouse"

    host = "irc.fc00.io"
    port = 6667
    tls = false
    debug = true
    storage = "sqlite3"

    reconnect = true
    ping = 30

    channels = [ "#lukevers", "#mouse", "#mice" ]

    plugins "javascript" {
        enabled = true
        folders = [ "contrib/scripts/javascript/" ]
        pattern = "*.js"
        events = [ "PRIVMSG" ]
    }

    store "sqlite3" {
        dsn = "/path/to/dbname.db"
    }

    store "mysql" {
        dsn = "user:password@/dbname?charset=utf8&parseTime=True"
    }

    store "postgres" {
        dsn = "host=myhost user=mouse dbname=mouse sslmode=disable password=mypassword"
    }

    store "mssql" {
        dsn = "server=localhost;user id=sa"
    }
}
