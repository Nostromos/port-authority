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

	// conn.Write() returns n, err where n is the number of bytes from your byteslice that weresuccessfully written to the connection
	status := "HTTP/1.1 200 OK\r\n\r\n"
	total := 0

	for total < len(status) {
		n, err := conn.Write([]byte(status))
		total += n
		if err != nil {
			fmt.Println("ERROR: timed out or some other problem")
			os.Exit(1)
		}
	}

	// Extract URL path from HTTP request
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
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

	fmt.Println(headerText)
	fmt.Println(lines)
}
