package main

import "container/list"

// Message
type Message struct {
	// incrementing id for messages
	Id int
	// the user that sent the message
	Username string `json:"username"`
	// the message content itself
	Content string `json:"content"`
}

// MessageStore
type MessageStore struct {
	msgs *list.List
}

func NewMessageStore() *MessageStore {
	l1 := list.New()
	store := MessageStore{l1}
	return &store
}

func (s *MessageStore) List() *list.List {
	return s.msgs
}

func (s *MessageStore) Add(m Message) {
	s.msgs.PushBack(m)
}
