package main

import (
	"fmt"
	"log"
	"net" // standard network package
	"strings"
)

func main() {
	// config
	port := 8000
	protocol := "tcp"

	// resolve TCP address
	addr, err := net.ResolveTCPAddr(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}

	// get TCP socket
	socket, err := net.ListenTCP(protocol, addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listen: ", socket.Addr().String())

	// keep listening
	for {
		// wait for connection
		conn, err := socket.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Connected by ", conn.RemoteAddr().String())

		// yield connection to concurrent process
		go handleConnection(conn)
	}
}

/*
func handleConnection(conn net.Conn) {
	// close connection when this function ends
	defer conn.Close()

	// write response
	conn.Write([]byte("Hello world."))
}
*/

func handleConnection(conn net.Conn) {
    // close connection when this function ends
    defer conn.Close()
 
    buf := make([]byte, 1024)
    conn.Read(buf)
 
    log.Printf("Request\n----------\n%s\n----------", string(buf))

	str := strings.Split(string(buf), " ")

	var res string
	switch str[1] {
		case "/hello":
			body := "Hello world."
			res = fmt.Sprintf("HTTP/1.1 200 OK\nContent-Length: %d\n\n%s", len([]byte(body)), body)
		case "/bye":
			body := "Good bye."
			res = fmt.Sprintf("HTTP/1.1 200 OK\nContent-Length: %d\n\n%s", len([]byte(body)), body)
		case "/hello.jp":
			body := "こんにちは"
			res = fmt.Sprintf("HTTP/1.1 200 OK\nContent-Type: text/plain; charset=UTF-8\nContent-Length: %d\n\n%s", len([]byte(body)), body)
		default:
			res = "HTTP/1.1 404 Not Found\n"
	}
 
    // write response
    //conn.Write([]byte("Hello world."))
	net.Conn.Write(conn, []byte(res))
}
