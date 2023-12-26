package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByNumber(int) (*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE if not exists accounts (
		id serial  PRIMARY KEY,
		first_name varchar(50),
		last_name varchar(50),
		number serial ,
		balance INT,
		created_at timestamp,
		encrypted_password varchar(256)
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `insert into accounts 
	(first_name, last_name, number,balance,created_at,encrypted_password) 
	values ($1, $2, $3, $4, $5,$6)`

	_, err := s.db.Query(query, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt, account.EncryptedPassword)

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM accounts WHERE id=$1", id)
	return err
}

func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts WHERE ID=$1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccounts(rows)
	}
	return nil, fmt.Errorf("Account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}
	accs := []*Account{}
	for rows.Next() {
		acc, err := ScanIntoAccounts(rows)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

func ScanIntoAccounts(rows *sql.Rows) (*Account, error) {
	acc := &Account{}
	if err := rows.Scan(
		&acc.ID,
		&acc.FirstName,
		&acc.LastName,
		&acc.Number,
		&acc.Balance,
		&acc.CreatedAt,
	); err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *PostgresStore) GetAccountByNumber(number int) (*Account, error) {
	rows, err := s.db.Query("select * from Account where number = $1", number)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccounts(rows)
	}
	return nil, fmt.Errorf("Account not found")
}
