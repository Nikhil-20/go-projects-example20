package task

type Store interface {
	Add(title string) (*Task, error)
	List() ([]Task, error)
	Complete(id int) error
	Delete(id int) error
	GetTaskDetails(id int) (*Task, error)
}
