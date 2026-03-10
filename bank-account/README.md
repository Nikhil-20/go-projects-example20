# Bank Account Management

A CLI-based bank account management system built in Go. Supports creating accounts, depositing and withdrawing funds, and viewing account details — all through an interactive terminal menu.

## Features

- **Create Account** with personal details (name, age, address, nationality, Aadhaar number, account type)
- **Deposit Amount** into an existing account
- **Withdraw Amount** with balance validation
- **Get Account Details** by account number
- Account data is persisted to JSON files on creation
- **Thread-safe** operations using `sync.Mutex`
- Input validation for account types (Savings / Current)

## Project Structure

```
bank-account/
├── main.go                          # Entry point with interactive CLI menu
├── go.mod
├── model/
│   └── bankmodel.go                 # Data models (AccountManager, AccountType)
├── accountmanager/
│   └── accountmanagement.go         # Business logic (CRUD operations)
└── bankstore/
    └── store.go                     # BankManagement interface definition
```

## Usage

```bash
cd bank-account
go run main.go
```

### Menu Options

```
Enter 1 to Add Account
Enter 2 to Deposit Amount
Enter 3 to Get Account Details
Enter 4 to Withdraw Amount
Enter 0 to Exit
```

## How It Works

1. On startup, an in-memory store is initialized via `InitializeStoreData()`.
2. Users interact through a CLI menu powered by `bufio.Scanner`.
3. Account creation generates a random account ID and saves account data as a JSON file.
4. Deposits and withdrawals update the in-memory store with mutex-based synchronization.

## Concepts Covered

- Structs and methods in Go
- Interfaces (`BankManagement`)
- Custom type validation (`AccountType`)
- `sync.Mutex` for concurrent access
- JSON marshalling and file I/O
- Interactive CLI with `bufio.Scanner`
