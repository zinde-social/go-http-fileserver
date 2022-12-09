package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	listenPort *int
	listenHost *string
)

func init() {
	listenPort = flag.Int("p", 0, "Listen to which port")
	listenHost = flag.String("l", "localhost", "Listen to which host")
}

func main() {

	// Parse arguments
	flag.Parse()

	// Get listen port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *listenHost, *listenPort))
	if err != nil {
		log.Fatalf("Failed to listen to target port: %d", listener.Addr().(*net.TCPAddr).Port)
	}

	// Serve current dir
	fs := http.FileServer(http.Dir("."))

	// Start HTTP server
	log.Printf("Starting HTTP server at http://%s", listener.Addr().String())
	err = http.Serve(listener, fs)
	if err != nil {
		log.Fatalf("Failed to start HTTP server with error: %v", err)
	}

}
