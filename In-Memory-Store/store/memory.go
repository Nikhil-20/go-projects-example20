package store

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	value      string
	expiration int64
}

type memorystore struct {
	mu    sync.Mutex
	data  map[string]item
	close chan struct{}
}

//constructor

func NewMemoryStore() *memorystore {

	fmt.Println("Initializing the new MemoryStore")
	data := make(map[string]item)
	close := make(chan struct{})
	s := memorystore{
		data:  data,
		close: close,
	}

	go s.StartCleanup()
	return &s
}

func (s *memorystore) Set(key string, value string, ttl time.Duration) {

	fmt.Println("Set method invoke")
	s.mu.Lock()
	defer s.mu.Unlock()

	var exp int64
	if ttl > 0 {
		exp = time.Now().Add(ttl).UnixNano()
	}

	s.data[key] = item{
		value:      value,
		expiration: exp,
	}

}

func (s *memorystore) Get(key string) (string, bool) {
	fmt.Println("Get method invoke")
	s.mu.Lock()
	defer s.mu.Unlock()

	it, exist := s.data[key]

	if !exist {
		return "", false
	}

	if it.expiration > 0 && time.Now().UnixNano() > it.expiration {
		return "", false
	}

	// var value string
	// if key != "" {
	// 	value = s.data[key].value
	// }

	return it.value, true
}

func (s *memorystore) Delete(key string) {
	fmt.Println("Delete method invoke")
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
	// _, exist := s.data[key]

	// if exist {
	// 	s.data[key] = item{
	// 		value:      "",
	// 		expiration: 0,
	// 	}
	// }
}

func (s *memorystore) Clear() {
	fmt.Println("In clear")
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[string]item)
}

func (s *memorystore) CLose() {
	close(s.close)
}

func (s *memorystore) StartCleanup() {
	fmt.Println("Start cleanup")
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("init cleanup case C")
			s.Cleanup()
		case <-s.close:
			fmt.Println("Close loop stopping ticker")
			ticker.Stop()
			return
		}
	}

}

func (s *memorystore) Cleanup() {
	fmt.Println("Cleanup of the data")
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano()
	for key, it := range s.data {
		if it.expiration > 0 && now > it.expiration {
			delete(s.data, key)
		}
	}
}
