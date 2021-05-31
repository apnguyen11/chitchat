package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var message string = ""

func main() {
	// Set routing rules
	http.HandleFunc("/messages/send", SetMessage)
	http.HandleFunc("/messages/receive", GetMessage)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func SetMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "CHIT CHAT!!!")
}
