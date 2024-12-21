package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "nian":
		fmt.Fprint(w, "No more room")
	default:
		fmt.Fprint(w, "Rock and Stone!")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlVsPlane(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hrmlVsPlain")
	//w.Header().Set("Content-Type", "text/plain") //to set html page to plain page "text/html" changes to "text/plain"
	fmt.Fprint(w, "<h1>Hello World!!!</h1>")
}

func timeoutThing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("timeoutThing")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Didn't timeout")
}

func MuffinPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MuffinPage request")
	//w.Header().Set("Content-Type", "text/plain") //to set html page to plain page "text/html" changes to "text/plain"
	fmt.Fprint(w, "<h1 style=\"background-color: #FF006E\">Hello World!!</h1>")
}

func main() {
	//http.HandleFunc("/", htmlVsPlane)
	//http.ListenAndServe("", nil)
	http.HandleFunc("/timeout", timeoutThing)

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
