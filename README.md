# Budgoteer

A simple webapp to track finances, tailored to my need, now with GO, templ and HTMX!

## Requirements

### MVP
- [ ] NO login
  - one "user" per application instance
- [ ] User can set one budget
  - one budget per month
  - user can set start date of budget (so it doesn't have to start on the first of the month)
- [ ] User can track expenses / reimbursements
  - These count against the budget

### Further
- [ ] profiles
   - user can switch between profiles
- [ ] multiple budgets per profile
- [ ] Transaction tagging
- [ ] data visualization
  - overview of historical data
  - analysis of data

## Technologies
- Go
- Templ
- HTMX
- SQLite

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
