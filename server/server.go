package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apnguyen11/chitchat/server/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var messages *MessageStore
var db *gorm.DB

func init() {
	messages = NewMessageStore()
}

func main() {
	log.Printf("starting server")

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.User{})

	// Set routing rules
	http.HandleFunc("/messages/send", SendMessage)
	http.HandleFunc("/messages/receive", GetMessage)
	http.HandleFunc("/register", UserRegister)
	http.HandleFunc("/login", UserLogin)

	//Use the default DefaultServeMux.
	err = http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
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

	enableCors(&w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var m *model.Message

	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(m)

	db.Create(m)
	// messages.Add(m)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var messages []model.Message

	enableCors(&w)

	db.Find(&messages)

	for _, msg := range messages {
		s := fmt.Sprintf("[%s] %s: %s \n", msg.Channel, msg.UserID, msg.Content)
		io.WriteString(w, s)
		// fmt.Println(e)
	}

	// for e := messages.Front(); e != nil; e = e.Next() {

	// 	msg := e.Value.(model.Message)
	// 	s := fmt.Sprintf("[%s] %s: %s \n", msg.Channel, msg.Username, msg.Content)
	// 	io.WriteString(w, s)
	// 	fmt.Println(e)

	// }

}

func UserRegister(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var registerRequest *model.UserRegisterRequest

	err = json.Unmarshal(body, &registerRequest)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(registerRequest)
	u := model.User{}
	u.Username = registerRequest.Username
	u.Password = registerRequest.Password

	db.Create(&u)
	// messages.Add(m)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var loginRequest *model.UserLoginRequest

	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		log.Println(err)
	}

	var user model.User

	//FIXME need to check error on return
	db.Where("username = ?", loginRequest.Username).First(&user)

	if user.Password == loginRequest.Password {
		log.Println("SUCCESS!!!")
	} else {
		log.Println("FAILL :(")
	}

	// messages.Add(m)
}
