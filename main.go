package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type GoservHandler struct {
	serv    http.Handler
	logging bool
}

func (gh *GoservHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if gh.logging {
		log.Printf("Serving: %s\n", req.URL.Path)
	}
	gh.serv.ServeHTTP(rw, req)
}

func main() {
	dir := "."
	dolog := flag.Bool("log", false, "enable logging")
	port := flag.Int("port", 8080, "specify port to listen on")
	flag.Parse()

	if len(flag.Args()) > 1 {
		dir = flag.Arg(0)
	}

	serv := http.FileServer(http.Dir(dir))
	gh := new(GoservHandler)
	gh.logging = *dolog
	gh.serv = serv

	for ; *port <= 65535; *port++ {
		fmt.Printf("Serving on port: %d\n", *port)
		http.ListenAndServe(fmt.Sprintf(":%d", *port), gh)
		fmt.Printf("Failed to serve on port: %d, retrying...\n", *port)
	}
	if *port > 65535 {
		fmt.Errorf("There is something seriously wrong with your computer.")
	}
}
