package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var messages *MessageStore

func init() {
	messages = NewMessageStore()
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

	var m Message

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	messages.Add(m)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {

	for e := messages.List().Front(); e != nil; e = e.Next() {

		msg := e.Value.(Message)
		s := fmt.Sprintf("%s: %s \n", msg.Username, msg.Content)
		io.WriteString(w, s)
		fmt.Println(e)

	}

}
