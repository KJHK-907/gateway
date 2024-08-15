package handlers

import (
	"context"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"gateway/models"
	"gateway/services"
)

func StartTCPServer(ctx context.Context, pool *models.Pool) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Println(err)
		listener.Close()
		return
	}
	ctxListener := newContextListener(ctx, listener)
	defer ctxListener.Close()
	for {
		select {
		case <-ctx.Done():
			log.Println("TCP server shutting down due to context cancellation")
			return
		default:
			conn, err := ctxListener.Accept()
			if err != nil {
				if ctx.Err() != nil {
					log.Println("TCP server shutting down due to context cancellation")
					return
				}
				log.Println("Error accepting connection:", err)
				continue
			}
			log.Println("Connected with Zetta RCS")
			go func() {
				handleConnection(ctx, conn, pool)
			}()
		}
	}
}

func handleConnection(ctx context.Context, conn net.Conn, pool *models.Pool) {
	defer conn.Close()
	var buffer strings.Builder
	for {
		select {
		case <-ctx.Done():
			log.Println("Connection handling cancelled")
			return
		default:
			buf := make([]byte, 4096)
			conn.SetReadDeadline(time.Now().Add(10 * time.Minute)) // Set a timeout of 10 minutes
			n, err := conn.Read(buf)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					log.Println("Read timeout, closing connection")
					return
				}
				if err == io.EOF {
					log.Println("Connection closed by Zetta RCS")
					return
				}
				log.Println("Read error:", err)
				return
			}
			if n > 0 {
				log.Println("Received metadata from Zetta RCS:")
				log.Println(string(buf[:n]))
				println("--------------------------")
				buffer.Reset()
				services.ProcessXml(buf[:n], &buffer, pool)
			}
		}
	}
}
