package main

import (
	"bytes"
	"strings"
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("ERROR: Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// Extract URL path from HTTP request
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	// write loop that populates and ensures all data gets written into buffer
	if err != nil {
		fmt.Println("ERROR: Something bad happened while readin the connection")
		os.Exit(1)
	}

	// find size of header and body if it exists
	headerEnd := bytes.Index(buf[:n], []byte("\r\n\r\n"))
	headerBytes := buf[:headerEnd]
	// bodyBytes := buf[headerEnd + 4 : n] // if body exists

	// headerLines := bytes.Split(headerBytes, []byte("\r\n"))

	headerText := string(headerBytes)
	lines := strings.Split(headerText, "\r\n")

	request := strings.Split(lines[0], " ");
	method := request[0]
	target := request[1]
	version := request[2]

	host := strings.Split(lines[1], " ");

	fmt.Println("Lines Length: ", len(lines))
	fmt.Println(method)
	fmt.Println(target)
	fmt.Println(version)
	fmt.Println(host)

	// fmt.Println(headerText)
	fmt.Println(lines[0])

	// conn.Write() returns n, err where n is the number of bytes from your byteslice that weresuccessfully written to the connection
	okStatus := "HTTP/1.1 200 OK\r\n\r\n"
	notFoundStatus := "HTTP/1.1 404 Not Found\r\n\r\n"
	total := 0

	for total < len(okStatus) {
		if (len(target) > 1) && (target != "/") {
			_, err := conn.Write([]byte(notFoundStatus))
			if err != nil {
				fmt.Println("ERROR: Cannot write to connection")
				os.Exit(1)
			}
		} else {
			_, err := conn.Write([]byte(okStatus))
			if err != nil {
				fmt.Println("ERROR: Cannot write to connection")
				os.Exit(1)
			}
		}
		// total += n
		os.Exit(0)
	}
}
