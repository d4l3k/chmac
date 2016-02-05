package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/d4l3k/chmac/mac"
)

var (
	hwinterface = flag.String("interface", "", "the network interface")
	resnet      = flag.Bool("resnet", false, "whether to poll every 60s for resnet caps")
)

func main() {
	flag.Parse()
	inter, err := net.InterfaceByName(*hwinterface)
	if err != nil {
		log.Fatal("Invalid interface")
	}
	if !*resnet {
		if err := mac.SetRandomMac(inter); err != nil {
			log.Fatal(err)
		}
	} else {
		timer := time.NewTimer(60 * time.Second)
		for range timer.C {
			resp, err := http.Get("http://ubcit.webi.it.ubc.ca/__shared/Pagelet5764.html")
			if err != nil {
				log.Fatal(err)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			if bytes.Contains(body, []byte("is currently in")) {
				log.Println("Randomizing mac address...")
				if err := mac.SetRandomMac(inter); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
