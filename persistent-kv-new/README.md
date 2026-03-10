# Persistent Key-Value Store

A persistent key-value store built in Go with a Write-Ahead Log (WAL) for durability and an HTTP API for interaction.

## Features

- **HTTP API** for setting and getting key-value pairs
- **Write-Ahead Log (WAL)** ensures data durability across restarts
- **Thread-safe** operations using `sync.Mutex`
- **Recovery** from WAL on startup
- Simple and lightweight — no external dependencies

## Project Structure

```
persistent-kv-new/
├── main.go          # Entry point — initializes WAL, store, and HTTP server
├── go.mod
└── store/
    ├── store.go     # In-memory store with WAL integration and recovery
    ├── wal.go       # Write-Ahead Log implementation (append-only JSON log)
    └── http.go      # HTTP handlers for /set and /get endpoints
```

## Usage

### Start the Server

```bash
cd persistent-kv-new
go run main.go
```

The server starts on port **8080**.

### API Endpoints

#### Set a key-value pair

```bash
curl "http://localhost:8080/set?key=name&value=Nikhil"
# Response: OK
```

#### Get a value by key

```bash
curl "http://localhost:8080/get?key=name"
# Response: {"value":"Nikhil"}
```

## How It Works

1. On startup, a WAL file (`wal.log`) is opened/created.
2. The `Store` is initialized with the WAL reference.
3. Every **Set** operation first appends a JSON log entry to the WAL, then updates the in-memory map.
4. The WAL uses `file.Sync()` to ensure entries are flushed to disk.
5. On recovery, WAL entries are replayed to reconstruct the in-memory state.

### WAL Log Format

Each line in `wal.log` is a JSON object:

```json
{"command":"SET","key":"name","value":"Nikhil"}
```

## Concepts Covered

- Write-Ahead Logging for persistence
- `sync.Mutex` for thread safety
- HTTP handlers with `net/http`
- JSON encoding/decoding
- File I/O with `os.OpenFile` (append mode)
- Separation of concerns (store, WAL, HTTP layers)
