package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type FileStore struct {
	mu     sync.Mutex
	file   string
	tasks  []Task
	nextID int
}

//constructor

func NewFilestore(filename string) (*FileStore, error) {

	fmt.Println("In NewFilestore")
	fs := FileStore{
		file: filename,
	}

	err := fs.load()
	if err != nil {
		return nil, err
	}

	return &fs, nil
}

func (fs *FileStore) load() error {
	fmt.Println("In load")
	data, err := os.ReadFile(fs.file)

	if err != nil {
		fmt.Println("after load if err:", err)
		if os.IsNotExist(err) {
			fmt.Println("err is not exist:")
			return nil
		}
		return err
	}

	fmt.Println("unmarshalling data")
	err = json.Unmarshal(data, &fs.tasks)
	if err != nil {
		return err
	}

	for _, ts := range fs.tasks {
		if ts.ID >= fs.nextID {
			fs.nextID = ts.ID + 1
		}
	}
	return nil
}

func (fs *FileStore) save() error {
	fmt.Println("In save")
	data, err := json.MarshalIndent(fs.tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.file, data, 0644)
}

func (fs *FileStore) Add(title string) (*Task, error) {

	fmt.Println("In Add")
	fs.mu.Lock()
	defer fs.mu.Unlock()

	task := Task{
		ID:        fs.nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	fs.tasks = append(fs.tasks, task)
	fs.nextID++

	if err := fs.save(); err != nil {
		return nil, err
	}

	return &task, nil
}

func (fs *FileStore) List() ([]Task, error) {
	fmt.Println("In list")
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return fs.tasks, nil
}

func (fs *FileStore) Complete(id int) error {
	fmt.Println("In Complete")
	fs.mu.Lock()
	defer fs.mu.Unlock()

	for i, _ := range fs.tasks {
		if fs.tasks[i].ID == id {
			fs.tasks[i].Completed = true
			return fs.save()
		}
	}

	return errors.New("task not found")
}

func (fs *FileStore) Delete(id int) error {
	fmt.Println("In Delete")
	fs.mu.Lock()
	defer fs.mu.Unlock()

	for i := range fs.tasks {
		if fs.tasks[i].ID == id {
			fs.tasks = append(fs.tasks[:i], fs.tasks[i+1:]...)
			return fs.save()
		}
	}

	return errors.New("task not found")
}
