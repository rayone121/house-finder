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
cd reBank
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

## Built With

- [Go](https://golang.org/) - The programming language used
- [PostgreSQL](https://www.postgresql.org/) - The database used

## Authors

- **Rayone** - *Initial work* - [rayone121](https://github.com/rayone121)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

