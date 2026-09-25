package p2p

// This func is 
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
