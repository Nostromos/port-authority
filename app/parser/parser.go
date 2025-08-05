package parser

import (
	"net/http"
	"bytes"
	"strings"

)

func ParseRequest(buf []byte, n int) {
	headerEnd := bytes.Index(buf[:n], []byte("\r\n\r\n"))
	headerBytes := buf[:headerEnd]

	headerText := string(headerBytes)
	lines := strings.Split(headerText, "\r\n")

	request := strings.Split(lines[0], " ");

	return &ParsedHTTPRequest{
		method: request[0],
		target: request[1],
		version: request[2],
		host: strings.Split(lines[1], " ")[1],
		headers: strings.Join(lines[2:], "\r\n"),
		body: "",
	}
}