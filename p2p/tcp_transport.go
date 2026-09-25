package p2p

import (
	"fmt"
	"net"
	"sync"
)

// This represents the remote node over the TCP established connection
type TCPPeer struct {
	// Underlying connection of the peer
	conn net.Conn

	// If we dial and receive a connection => outbound == true
	// If we accept and retrieve a conn => outbound == false
	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NOPHandshakFunc(any)
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands: NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

type Temp struct {}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s", err)
		}

		fmt.Printf("New incoming conntion %+v\n", conn)

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(conn); err != nil {

	}

	msg := &Temp{}

	// Read loop
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}

	fmt.Printf("New incoming connection %+v\n", peer)
}
