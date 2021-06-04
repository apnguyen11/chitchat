package main

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	store := NewMessageStore()
	store.Add(Message{0, "buddy", "hello"})
	store.Add(Message{1, "joe", "sawp"})

	// print messages
	for e := store.List().Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
}
