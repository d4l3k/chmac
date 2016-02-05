package main

import (
	"flag"
	"log"
	"net"

	"github.com/d4l3k/chmac/mac"
)

var hwinterface = flag.String("interface", "", "the network interface")

func main() {
	flag.Parse()
	inter, err := net.InterfaceByName(*hwinterface)
	if err != nil {
		log.Fatal("Invalid interface")
	}
	if err := mac.SetRandomMac(inter); err != nil {
		log.Fatal(err)
	}
}
