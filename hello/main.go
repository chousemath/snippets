package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var input string
	fmt.Printf("Enter your name: ")
	fmt.Scanln(&input)
	fmt.Printf("Hello %s!\n", input)

	port := flag.String("p", "8000", "port")
	dir := flag.String("d", ".", "dir")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Hi %s, I am serving %s on http port: %s\n", input, *dir, *port)
	// no handler function in this case
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
