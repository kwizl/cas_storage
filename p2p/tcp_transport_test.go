package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"

	tcpOpts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       GOBDecoder{},
	}

	tr := NewTCPTransport(tcpOpts)

	// Port
	assert.Equal(t, tr.ListenAddr, listenAddr)

	// Server
	assert.Nil(t, tr.ListenAndAccept())

}
