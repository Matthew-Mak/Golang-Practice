package main

import (
	"fmt"
	"net/http"
)

type MD struct {
	Name string
}

var cyn = MD{"Cyn"}

func MuffinPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		//url(w, r)
	default:
		fmt.Println("MuffinPage request")
		//w.Header().Set("Content-Type", "text/plain") //to set html page to plain page "text/html" changes to "text/plain"
		// w.Write([]byte("Hello World"))
		fmt.Fprint(w, "Hi xdx")
	}
}

func main() {
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxMuffinTime http.ServeMux
	server.Handler = &muxMuffinTime
	muxMuffinTime.HandleFunc("/", MuffinPage)

	server.ListenAndServe()
}
