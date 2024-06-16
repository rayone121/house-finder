# reBank

reBank is a simple banking application written in Go. It allows users to create accounts, transfer money between accounts, and view account balances.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go programming language
- PostgreSQL

### Installing

1. Clone the repository
```bash
git clone https://github.com/rayone121/reBank.git
```

2. Navigate to the project directory
```bash
cd ./reBank
```

3. Install the necessary Go packages
```bash
go mod tidy
```

4. Run the application
```bash
make pg
make run
```
# reBank API Documentation

## Overview
reBank is a simple banking application that allows users to create accounts, transfer money between accounts, and view account balances.

## Endpoints

### `POST /login`
Handles user login.

### `GET /account`
Returns the account details of the logged-in user.

### `POST /account`
Creates a new account.

### `GET /account/{id}`
Returns the account details of a specific user by ID.

### `DELETE /account/{id}`
Deletes a specific user account by ID.

### `POST /transfer`
Handles money transfer requests between accounts.

## Account Structure
Accounts are represented as follows:

```go
type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	UserName          string    `json:"userName"`
	EncryptedPassword string    `json:"_"`
	Balance           int64     `json:"balance"`
	Number            int64     `json:"number"`
	CreatedAt         time.Time `json:"createdAt"`
}
```

## Running the Server
To run the server, execute the `main` function in `main.go`. You can seed the database with test data by passing the `--seed true` flag:

```sh
go run main.go --seed true
```

## Built With

- [Go](https://golang.org/) - The programming language used
- [PostgreSQL](https://www.postgresql.org/) - The database used

## Authors

- **Rayone** - *Initial work* - [rayone121](https://github.com/rayone121)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

