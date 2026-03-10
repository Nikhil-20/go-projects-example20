# In-Memory Store

A thread-safe in-memory key-value store built in Go with TTL (Time-To-Live) support and automatic cleanup of expired entries.

## Features

- **Set** a key-value pair with an optional TTL duration
- **Get** a value by key (returns empty if expired or missing)
- **Delete** a specific key
- **Clear** all stored data
- **Automatic cleanup** of expired keys via a background goroutine
- **Thread-safe** operations using `sync.Mutex`

## Project Structure

```
In-Memory-Store/
├── main.go            # Entry point demonstrating usage
├── go.mod
└── store/
    ├── store.go       # Storage interface definition
    └── memory.go      # In-memory implementation with TTL & cleanup
```

## How It Works

1. `NewMemoryStore()` initializes the store and starts a background cleanup goroutine that runs every second.
2. `Set(key, value, ttl)` stores a key-value pair with an expiration timestamp derived from the TTL.
3. `Get(key)` retrieves the value only if the key exists and hasn't expired.
4. The background `StartCleanup()` goroutine periodically removes expired entries from the map.

## Usage

```bash
cd In-Memory-Store
go run main.go
```

### Example

```go
s := store.NewMemoryStore()

s.Set("name", "nn", 5*time.Second)

val, found := s.Get("name")  // found = true
time.Sleep(6 * time.Second)
val, found = s.Get("name")   // found = false (expired)
```

## Concepts Covered

- Interfaces in Go
- Goroutines and background workers
- `sync.Mutex` for concurrent access
- TTL-based expiration logic
- Channel-based shutdown signaling
