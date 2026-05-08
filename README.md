# Go Double Entry Ledger

A simple double-entry accounting ledger API built with Go.

## Features

- Create accounts
- Get account by ID
- Create balanced transactions
- Apply ledger entries to account balances
- In-memory storage
- Domain-driven design
- Unit tests

## Business Rules

- Every transaction must be balanced: total debits == total credits.
- Every transaction must have at least one entry.
- Each entry must have amount > 0 and a valid direction (`debit` or `credit`).
- When applying an entry to an account:
  - Same direction as account: increase balance
  - Different direction: decrease balance
- Money values are represented as int64 in cents (no floats).

## Architecture

```
cmd/api                      # Application entry point
internal/domain              # Business entities and rules
internal/service             # Orchestration and use cases
internal/repository          # Repository interfaces
internal/repository/memory   # In-memory repository implementations
internal/handler/http        # HTTP handlers (API layer)
```

## How to Run (Locally)

```sh
go run cmd/api/main.go
```

Server runs at:  
http://localhost:5000

## How to Run with Docker

Build and run the application using Docker Compose:

```sh
docker compose up --build
```

The API will be available at:
http://localhost:5000

> **Note:** Ensure Docker is installed and the Docker daemon is running. In WSL2, start Docker Desktop on Windows.

## How to Test

```sh
go test ./...
```

## API Examples

### Create Account
```sh
curl -X POST http://localhost:5000/accounts \
  -H "Content-Type: application/json" \
  -d '{"id":"acc-1","name":"Cash","direction":"debit","balance":0}'
```

### Get Account
```sh
curl http://localhost:5000/accounts/acc-1
```

### Create Transaction
```sh
curl -X POST http://localhost:5000/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "id":"tx-1",
    "name":"Initial movement",
    "entries":[
      {
        "id":"entry-1",
        "account_id":"acc-1",
        "direction":"debit",
        "amount":100
      },
      {
        "id":"entry-2",
        "account_id":"acc-2",
        "direction":"credit",
        "amount":100
      }
    ]
  }'
```

### Get Transaction
```sh
curl http://localhost:5000/transactions/tx-1
```
