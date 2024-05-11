package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rayone121/house-finder/backend/types"
)

type Storage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccountByID(int) (*types.Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "postgres://postgres:1234@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		username varchar(32), 
		first_name varchar(32),
		last_name varchar(32), 
		number serial,
		balance integer,
		create_at timestamp
		)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) Close() {
	s.db.Close()
}

func (s *PostgressStore) CreateAccount(a *types.Account) error {
	return nil
}

func (s *PostgressStore) UpdateAccount(a *types.Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgressStore) GetAccountByID(id int) (*types.Account, error) {
	return nil, nil
}
