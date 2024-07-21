package handlers

import (
	"fmt"
	"log"
	"net"
	"strings"

	"gateway/services"
)

func StartTCPServer() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	conn, err := listener.Accept()
	println("Connected with Zetta RCS")
	if err != nil {
		fmt.Println(err)
		return
	}
	go handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buffer strings.Builder
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		if n > 0 {
			println("Received metadata from Zetta RCS:")
			log.Println(string(buf[:n]))
			println("--------------------------")
			services.ProcessXml(buf[:n], &buffer)
		}
	}
}
