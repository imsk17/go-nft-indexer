package listeners

// Listener represents an event Listener.
type Listener interface {
	Listen()
	Close()
}
