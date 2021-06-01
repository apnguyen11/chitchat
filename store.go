package main

import "container/list"

// Message
type Message struct {
	// incrementing id for messages
	id int
	// the user that sent the message
	username string
	// the message content itself
	content string
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
