package main

import (
	"In-Memory-Store/store"
	"fmt"
	"time"
)

func main() {
	s := store.NewMemoryStore()

	s.Set("name", "nn", 5*time.Second)

	val, flag := s.Get("name")
	if flag {
		fmt.Println("Printing name:", val)
	}

	time.Sleep(6 * time.Second)

	val, flag = s.Get("name")

	if !flag {
		fmt.Println("key expired")
	}
	//s.Set("name")
}
