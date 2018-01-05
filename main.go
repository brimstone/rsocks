package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/yamux"
)

var session *yamux.Session

func main() {

	listen := flag.String("listen", "", "listen port for receiver address:port")
	socks := flag.String("socks", "127.0.0.1:1080", "socks address:port")
	connect := flag.String("connect", "", "connect address:port")
	version := flag.Bool("version", false, "version information")
	flag.Usage = func() {
		fmt.Println("rsocks - reverse socks5 server/client")
		fmt.Println("https://github.com/brimstone/rsocks")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("1) Start rsocks -listen :8080 -socks 127.0.0.1:1080 on the client.")
		fmt.Println("2) Start rsocks -connect client:8080 on the server.")
		fmt.Println("3) Connect to 127.0.0.1:1080 on the client with any socks5 client.")
		fmt.Println("4) Enjoy. :]")
	}

	flag.Parse()

	if *version {
		fmt.Println("rsocks - reverse socks5 server/client")
		fmt.Println("https://github.com/brimstone/rsocks")
		os.Exit(0)
	}

	if *listen != "" {
		log.Println("Starting to listen for clients")
		go listenForSocks(*listen)
		log.Fatal(listenForClients(*socks))
	}

	if *connect != "" {
		log.Println("Connecting to the far end")
		log.Fatal(connectForSocks(*connect))
	}

	fmt.Fprintf(os.Stderr, "You must specify a listen port or a connect address")
	os.Exit(1)
}
