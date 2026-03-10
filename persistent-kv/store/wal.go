package store

import (
	"encoding/json"
	"os"
	"sync"
)

type Wal struct {
	mu   sync.Mutex
	file *os.File
}

type LogEntry struct {
	Command string `json:"command"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

func NewWal(path string) (*Wal, error) {

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &Wal{file: f}, nil

}

func (w *Wal) Append(cmd, key, value string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := LogEntry{
		Command: cmd,
		Key:     key,
		Value:   value,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = w.file.Write(append(data, '\n'))
	if err != nil {
		return err
	}

	return w.file.Sync()
}
