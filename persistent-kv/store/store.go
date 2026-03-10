package store

import (
	"fmt"
	"sync"
)

type Store struct {
	mu   sync.Mutex
	data map[string]string
	wal  *Wal
}

func NewStore(wal *Wal) *Store {
	return &Store{
		data: make(map[string]string),
		wal:  wal,
	}
}

func (s *Store) Set(key, value string) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.wal.Append("SET", key, value); err != nil {
		return err
	}
	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.data[key]
	if !ok {
		fmt.Println("Error in getting key")
	}
	return val, ok
}

func (s *Store) Recover(entries []LogEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, e := range entries {
		if e.Command == "SET" {
			s.data[e.Key] = e.Value
		}
	}
}
