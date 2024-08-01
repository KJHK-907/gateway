package handlers

import (
	"io"
	"log"
	"net"
	"strings"
	"time"

	"gateway/models"
	"gateway/services"
)

func StartTCPServer(pool *models.Pool) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		log.Println("Connected with Zetta RCS")
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, pool)
	}
}

func handleConnection(conn net.Conn, pool *models.Pool) {
	defer conn.Close()
	var buffer strings.Builder
	for {
		buf := make([]byte, 1024)
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
