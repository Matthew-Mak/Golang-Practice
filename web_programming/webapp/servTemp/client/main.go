package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	//plain old URL - http://localhost
	response, err := http.Get("http://localhost:80")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	//URL key-value form - http://localhost/url?name=Cyn
	response, err = http.PostForm(
		"http://localhost/url",
		url.Values{"name": {"Cyn"}})
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// http://localhost/body
	// with body in json: {"name": "Cyn"}
	type MD struct {
		Name string
	}
	cyn := MD{"Cyn"}
	cynJson, _ := json.Marshal(cyn)
	response, err = http.Post(
		"http://localhost/body",
		"application/json",
		bytes.NewBuffer(cynJson))
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
