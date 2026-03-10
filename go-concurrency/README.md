# Go Concurrency Examples

A collection of Go programs demonstrating concurrency patterns using goroutines, channels, WaitGroups, and context.

## Projects

### 1. DNS Resolver

A concurrent DNS resolver that resolves multiple domain names in parallel using a worker pool pattern.

**Features:**
- Worker pool with configurable concurrency (default: 3 workers)
- Rate limiting using `time.Tick`
- Retry logic with exponential backoff (up to 3 attempts)
- Context-based timeout (5 seconds)
- Job queue via channels

**Usage:**
```bash
cd dns-resolver
go run main.go
```

**Example Output:**
```
google.com -> [142.250.x.x ...]
github.com -> [140.82.x.x ...]
golang.org -> [142.250.x.x ...]
```

---

### 2. URL Fetcher

A concurrent HTTP URL fetcher that checks the status of multiple URLs in parallel.

**Features:**
- Concurrent HTTP GET requests using goroutines
- Results collected via channels
- Error handling for unreachable URLs
- `sync.WaitGroup` for synchronization

**Usage:**
```bash
cd url-fetcher
go run main.go
```

**Example Output:**
```
Fetched the url https://google.com with status 200
Fetched the url https://github.com with status 200
Error : http://youtube1.com <error details>
```

## Project Structure

```
go-concurrency/
├── README.md
├── dns-resolver/
│   ├── go.mod
│   └── main.go
└── url-fetcher/
    ├── go.mod
    └── main.go
```

## Concepts Covered

- Goroutines and `sync.WaitGroup`
- Buffered and unbuffered channels
- Worker pool pattern
- Rate limiting with `time.Tick`
- `context.WithTimeout` for cancellation
- Retry logic with backoff
- Channel direction (`chan<-`, `<-chan`)
