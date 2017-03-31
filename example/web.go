package main

import (
	_ "github.com/qodrorid/godaemon"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Assalamu'alaikum, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":3030", mux))
}