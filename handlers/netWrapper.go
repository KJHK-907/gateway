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
	type connError struct {
		conn net.Conn
		err  error
	}

	ch := make(chan connError, 1)
	go func() {
		conn, err := l.Listener.Accept()
		ch <- connError{conn, err}
	}()

	select {
	case <-l.ctx.Done():
		return nil, l.ctx.Err()
	case ce := <-ch:
		return ce.conn, ce.err
	}
}

func (l *contextListener) Close() error {
	return l.Listener.Close()
}

func newContextListener(ctx context.Context, listener net.Listener) net.Listener {
	return &contextListener{Listener: listener, ctx: ctx}
}
