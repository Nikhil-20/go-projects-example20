package store

import "sync"

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

	err := s.wal.Append("SET", key, value)
	if err != nil {
		return err
	}
	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.data[key]

	if !ok {
		return "", false
	}

	return v, ok
}

func (s *Store) Recover(entries []LogEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, entry := range entries {
		if entry.Command == "SET" {
			s.data[entry.Key] = entry.Value
		}
	}
}
