package handlers

import (
	"context"
	"net"
)

type contextListener struct {
	net.Listener
	ctx context.Context
}

func (l *contextListener) Accept() (net.Conn, error) {
	for {
		conn, err := l.Listener.Accept()
		select {
		case <-l.ctx.Done():
			return nil, l.ctx.Err()
		default:
			if err != nil {
				return nil, err
			}
			return conn, nil
		}
	}
}

func newContextListener(ctx context.Context, listener net.Listener) net.Listener {
	return &contextListener{Listener: listener, ctx: ctx}
}
