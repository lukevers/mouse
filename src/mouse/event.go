package mouse

// An Event represents an IRC event
type Event struct {
	Channel string
	Command string
	Host    string
	Message string
	Nick    string
	User    string
}
