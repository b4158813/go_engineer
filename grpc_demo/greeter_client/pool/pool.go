/*
	grpc 客户端连接池 demo
*/

package pool

import (
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ClientPool struct {
	pool sync.Pool
}

func GetPool(target string, opts ...grpc.DialOption) (*ClientPool, error) {
	return &ClientPool{
		pool: sync.Pool{
			New: func() any {
				conn, err := grpc.Dial(target, opts...)
				if err != nil {
					log.Fatalln(err)
				}
				return conn
			},
		},
	}, nil
}

func (c *ClientPool) Get() *grpc.ClientConn {
	conn := c.pool.Get().(*grpc.ClientConn)
	state := conn.GetState()
	if state == connectivity.TransientFailure || state == connectivity.Shutdown {
		conn.Close()
		conn = c.pool.New().(*grpc.ClientConn)
	}
	return conn
}

func (c *ClientPool) Put(conn *grpc.ClientConn) {
	state := conn.GetState()
	if state == connectivity.TransientFailure || state == connectivity.Shutdown {
		conn.Close()
		return
	}
	c.pool.Put(conn)
}
