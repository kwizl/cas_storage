package p2p

// Peer is an interface that represents the remote node in the network.
type Peer interface  {}

// Transport is anything that handles communication between nodes in the network
// This can be in the form of TCP, UDP & Webscokets
type Transport interface {
	ListenAndAccept() error
}
