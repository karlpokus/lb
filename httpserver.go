package main

import (
	"log"
	"net/http"
	"os"
)

var port = os.Args[1]

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("foo", "baloo")
	w.Write([]byte("hi from " + port))
}

func main() {
	http.HandleFunc("/foo", foo)
	log.Printf("http server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
