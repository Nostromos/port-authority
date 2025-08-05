package main

import (
	// "bytes"
	"fmt"
	"io"
	"net"
	"os"
	// "strings"

	"github.com/codecrafters-io/http-server-starter-go/logger"
	"github.com/codecrafters-io/http-server-starter-go/parser"
	"github.com/codecrafters-io/http-server-starter-go/response"
	"github.com/codecrafters-io/http-server-starter-go/router"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR: Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close(); // close the connection after this function returns
	/* On connection, createa  buffer and read into it */
	buf, err := io.ReadAll(conn)

	req, err := parser.ParseRequest(buf)
	fmt.Println(req)
	if err != nil {
		conn.Write(response.BuildResponse(WithStatus(400)))
		fmt.Printf("ERROR: Something bad happened while readin the connection: %v", err)
		os.Exit(1)
	}
}

	// // conn.Write() returns n, err where n is the number of bytes from your byteslice that weresuccessfully written to the connection
	// okStatus := "HTTP/1.1 200 OK\r\n\r\n"
	// notFoundStatus := "HTTP/1.1 404 Not Found\r\n\r\n"
	// var CRLF = "\r\n\r\n"
	// total := 0

	// for total < len(okStatus) {
	// 	if (len(target) > 1) && (target != "/") {
	// 		_, err := conn.Write([]byte(notFoundStatus))
	// 		if err != nil {
	// 			fmt.Println("ERROR: Cannot write to connection")
	// 			os.Exit(1)
	// 		}
	// 	} else {
	// 		_, err := conn.Write([]byte(okStatus))
	// 		if err != nil {
	// 			fmt.Println("ERROR: Cannot write to connection")
	// 			os.Exit(1)
	// 		}
	// 	}
	// 	// total += n
	// 	os.Exit(0)
	// }