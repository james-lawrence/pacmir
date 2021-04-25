package muxer_test

import (
	"context"
	"net"
	"sync/atomic"
	"testing"

	"github.com/james-lawrence/pacmir/internal/testingx"
	. "github.com/james-lawrence/pacmir/muxer"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMuxer(t *testing.T) {
	g := testingx.Init(t)

	g.Describe("Dial and accept", func() {
		g.It("should be able to multiplex", func() {
			c1 := int64(0)
			c2 := int64(0)
			counter := func(l net.Listener, c *int64) {
				for {
					_, err := l.Accept()
					if err != nil {
						return
					}
					atomic.AddInt64(c, 1)
				}
			}

			m := New()
			l, err := net.Listen("tcp", ":0")
			go func() {
				Listen(context.Background(), m, l)
			}()
			require.Nil(t, err)
			defer l.Close()

			l1, err := m.Bind("proto1", l.Addr())
			require.Nil(t, err)
			go counter(l1, &c1)
			l2, err := m.Bind("proto2", l.Addr())
			require.Nil(t, err)
			go counter(l2, &c2)

			d1 := NewDialer("proto1", &net.Dialer{})

			conn, err := d1.DialContext(context.Background(), "tcp", l.Addr().String())
			require.Nil(t, err)
			require.Nil(t, conn.Close())
			require.Equal(t, atomic.LoadInt64(&c1), int64(1))
			require.Equal(t, atomic.LoadInt64(&c2), int64(0))

			d2 := NewDialer("proto2", &net.Dialer{})

			conn, err = d2.DialContext(context.Background(), "tcp", l.Addr().String())
			require.Nil(t, err)
			require.Nil(t, conn.Close())
			require.Equal(t, atomic.LoadInt64(&c1), int64(1))
			require.Equal(t, atomic.LoadInt64(&c2), int64(1))
		})
	})

	g.Describe("Requested", func() {
		g.Describe("should encode to a fixed size (20)", func() {
			test := func(name string) {
				protocol := Proto(name)
				encoded, err := proto.Marshal(&Requested{
					Version:  1,
					Protocol: protocol[:],
				})
				require.Nil(t, err)
				require.Equal(t, len(encoded), 20)
			}

			g.It("example 1 - default protocol", func() { test("") })
			g.It("example 2 - long name", func() { test("45b3058c-9ec4-41cc-bd4c-c74ca2abdea2") })
		})
	})

	g.Describe("Accepted", func() {
		g.Describe("should encode to a fixed size(22)", func() {
			test := func(name string, code AcceptedError) {
				protocol := Proto(name)
				encoded, err := proto.Marshal(&Accepted{
					Version:  1,
					Code:     code,
					Protocol: protocol[:],
				})
				require.Nil(t, err)
				require.Equal(t, len(encoded), 22)
			}

			g.It("example 1 - default protocol", func() { test("", Accepted_None) })
			g.It("example 2 - long name", func() { test("45b3058c-9ec4-41cc-bd4c-c74ca2abdea2", Accepted_None) })
			g.It("example 3 - client error", func() { test("45b3058c-9ec4-41cc-bd4c-c74ca2abdea2", Accepted_ClientError) })
			g.It("example 4 - server error", func() { test("45b3058c-9ec4-41cc-bd4c-c74ca2abdea2", Accepted_ServerError) })
		})
	})
}
