package main

import (
	"golangweb/handle"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/main", handle.Mainhandle)
	mux.HandleFunc("/", handle.Homehandle)
	filehandle := http.FileServer(http.Dir("views"))
	mux.Handle("/static/", http.StripPrefix("/static/", filehandle))
	log.Print("Port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Print(err.Error())
	}
}
