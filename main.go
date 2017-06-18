package main

import (
	"fmt"
	"log"
	"net"
)

const (
	message       = "Hello World!"
	numIterations = 5
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		log.Fatalf("could not resolve UDP address: %v", err)
	}

	fmt.Println(addr)

	server, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("could not start UDP listener on %v: %v", addr, err)
	}
	defer server.Close()

	localAddr := server.LocalAddr()
	rAddr, err := net.ResolveUDPAddr(localAddr.Network(), localAddr.String())
	if err != nil {
		log.Fatalf("could not resolve local address %v: %v", localAddr.String(), err)
	}

	client, err := net.DialUDP("udp", nil, rAddr)
	if err != nil {
		log.Fatalf("could not create UDP client: %v", err)
	}
	defer client.Close()

	go func() {
		for i := 0; i < numIterations; i++ {
			if _, err := client.Write([]byte(message)); err != nil {
				log.Fatalf("could not write UDP message: %v", err)
			}
			log.Printf("UDP client sent the following message: '%v'", message)
		}
	}()

	for i := 0; i < numIterations; i++ {
		buf := make([]byte, len(message)/2)
		n, err := server.Read(buf)
		if err != nil {
			log.Fatalf("server failed to read message: %v", err)
		}
		log.Printf("server read the message: %v", string(buf[0:n]))
	}
}
