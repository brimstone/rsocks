package main

import (
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
)

func connectForSocks(address string) error {
	server, err := socks5.New(&socks5.Config{})
	if err != nil {
		return err
	}

	var conn net.Conn
	log.Println("Connecting to far end")
	conn, err = net.Dial("tcp", address)
	if err != nil {
		return err
	}

	log.Println("Starting server")
	session, err = yamux.Server(conn, nil)
	if err != nil {
		return err
	}

	for {
		stream, err := session.Accept()
		log.Println("Acceping stream")
		if err != nil {
			return err
		}
		log.Println("Passing off to socks5")
		go func() {
			err = server.ServeConn(stream)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}
