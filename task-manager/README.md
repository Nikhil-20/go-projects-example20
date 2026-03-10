# Task Manager

A CLI-based task manager built in Go that stores tasks in a JSON file. Supports adding, listing, completing, and deleting tasks.

## Features

- **Add** tasks with a title
- **List** all tasks with their completion status
- **Complete** a task by ID
- **Delete** a task by ID
- Tasks are persisted to a JSON file (`Newtask.json`)
- Auto-incrementing task IDs
- **Thread-safe** operations using `sync.Mutex`

## Project Structure

```
task-manager/
├── main.go            # Entry point with CLI argument parsing
├── go.mod
├── Newtask.json       # Task data file (auto-created)
└── task/
    ├── model.go       # Task struct definition
    ├── store.go       # Store interface
    └── file_store.go  # File-based store implementation
```

## Usage

```bash
cd task-manager
```

### Add a Task

```bash
go run main.go add "Buy groceries"
# Output: Added task with ID: 0
```

### List All Tasks

```bash
go run main.go list
# Output:
# [0] Buy groceries (completed: false)
# [1] Write README (completed: false)
```

### Complete a Task

```bash
go run main.go complete 0
```

### Delete a Task

```bash
go run main.go delete 1
```

## How It Works

1. On startup, `NewFilestore("Newtask.json")` loads existing tasks from the JSON file (or starts fresh if the file doesn't exist).
2. CLI commands are parsed from `os.Args`.
3. Each operation (add/list/complete/delete) is handled via the `FileStore` which implements the `Store` interface.
4. After every write operation, the task list is saved back to the JSON file using `json.MarshalIndent`.

### Data Format (`Newtask.json`)

```json
[
  {
    "id": 0,
    "title": "Buy groceries",
    "completed": false,
    "created_at": "2026-03-10T10:00:00Z"
  }
]
```

## Concepts Covered

- Interfaces in Go (`Store`)
- File-based persistence with JSON
- CLI argument parsing with `os.Args`
- `sync.Mutex` for thread safety
- Struct methods and pointer receivers
- Error handling patterns
