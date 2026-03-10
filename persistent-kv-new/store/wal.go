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
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return nil, err
	}

	return &Wal{file: file}, nil
}

func (wal *Wal) Append(command, key, value string) error {
	wal.mu.Lock()
	defer wal.mu.Unlock()

	entry := LogEntry{
		Command: command,
		Key:     key,
		Value:   value,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = wal.file.Write(append(data, '\n'))
	if err != nil {
		return err
	}

	return wal.file.Sync()
}
