package main

import "container/list"
import "github.com/apnguyen11/chitchat/server/model"

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

func (s *MessageStore) Add(m model.Message) {
	s.msgs.PushBack(m)
}
