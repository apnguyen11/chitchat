package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

var messages *MessageStore

func init() {
	messages =  NewMessageStore()
}

func main() {
	// Set routing rules
	http.HandleFunc("/messages/send", SendMessage)
	http.HandleFunc("/messages/receive", GetMessage)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	messages.Add(Message{0, "John", string(body[:])})

	w.Write(body)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {

	for e := messages.List().Front(); e != nil; e = e.Next() {

		msg := e.Value.(Message);
		s := fmt.Sprintf("%s: %s \n", msg.username, msg.content)
		io.WriteString(w, s)
		fmt.Println(e)
	
		
	}
	
}
