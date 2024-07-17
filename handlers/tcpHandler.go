package handlers

import (
	"fmt"
	"net"

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
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if n > 0 {
			currentMetadata, err := services.ParseXml(buf[:n])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%+v\n", currentMetadata)
		}
		// fmt.Println(string(buf))
	}
}
