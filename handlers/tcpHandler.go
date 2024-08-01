package handlers

import (
	"io"
	"log"
	"net"
	"strings"

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
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by Zetta RCS")
				return
			}
			log.Println(err)
			continue
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
