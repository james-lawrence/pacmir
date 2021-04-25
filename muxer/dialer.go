package muxer

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"sync/atomic"

	"github.com/pkg/errors"
)

type dialer interface {
	DialContext(ctx context.Context, network string, address string) (net.Conn, error)
}

var i = new(int64)

// Newdialer net.Dialer for the given protocol.
func NewDialer(protocol string, d dialer) Dialer {
	return Dialer{
		id:       atomic.AddInt64(i, 1),
		protocol: protocol,
		digest:   Proto(protocol),
		d:        d,
	}
}

// Dialer implements the net.Dialer interface
type Dialer struct {
	id       int64
	protocol string
	digest   Protocol
	d        dialer
}

func (t Dialer) Dial(network string, address string) (conn net.Conn, err error) {
	return t.DialContext(context.Background(), network, address)
}

func (t Dialer) DialContext(ctx context.Context, network string, address string) (conn net.Conn, err error) {
	// log.Printf("muxer.DialContext initiated: %T %s %s %s\n", t.d, t.protocol, network, address)
	// defer log.Printf("muxer.DialContext completed: %T %s %s %s\n", t.d, t.protocol, network, address)

	if conn, err = t.d.DialContext(ctx, network, address); err != nil {
		return conn, errors.Wrapf(err, "muxer.DialContext failed: %s %s://%s", t.protocol, network, address)
	}

	if tlsconn, ok := conn.(*tls.Conn); ok {
		if err := tlsconn.Handshake(); err != nil {
			conn.Close()
			return nil, errors.Wrap(err, "tls handshake failed")
		}

		s := tlsconn.ConnectionState()
		if s.NegotiatedProtocol != "bw.mux" {
			return conn, nil
		}
	}

	if err = handshakeOutbound(t.digest[:], conn); err != nil {
		log.Println("muxer.DialContext handshakeOutbound", t.protocol, network, address, err)
		conn.Close()
		return nil, err
	}

	return conn, nil
}
