package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	flag.Usage = func() {
		fmt.Printf("servethem v0.1\n-d directory to host (default current dir)\n-p port to serve on (default 8100)\n")
	}
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "The directory to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))
	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Printf("Others can reach you on:\nhttp://%s:%s\n", getOutboundIP(), *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
