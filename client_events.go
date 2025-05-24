package steam

// When this event is emitted by the Client, the connection is automatically closed.
// This may be caused by a network error, for example.
type FatalErrorEvent error

type ConnectedEvent struct{}

type DisconnectedEvent struct{}
