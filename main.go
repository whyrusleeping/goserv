package main

import (
	"net/http"
	"os"
	"fmt"
)

func main() {
	port := 8080
	dir := "."

	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	serv := http.FileServer(http.Dir(dir))

	for port := 8080; port <= 65535; port++ {
		http.ListenAndServe(fmt.Sprintf(":%d", port), serv)
	}
	if port > 65535 {
		fmt.Errorf("There is something seriously wrong with your computer.")
	}
}
