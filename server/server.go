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
	err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func SendMessage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w);
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var m Message

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(m)

	messages.Add(m)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {

	enableCors(&w);

	for e := messages.List().Front(); e != nil; e = e.Next() {

		msg := e.Value.(Message)
		s := fmt.Sprintf("[%s] %s: %s \n", msg.Channel, msg.Username, msg.Content)
		io.WriteString(w, s)
		fmt.Println(e)

	}

}
