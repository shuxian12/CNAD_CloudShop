package repository

import (
	"database/sql"
	"fmt"

	"CNAD_CloudShop/src/domain"
)

type SQLiteUserRepo struct {
	db *sql.DB
}

func NewSQLiteUserRepo(db *sql.DB) UserRepo {
	return &SQLiteUserRepo{db: db}
}

func (repo *SQLiteUserRepo) Create(user *domain.User) error {
	query := "INSERT INTO users (username) VALUES (?)"
	_, err := repo.db.Exec(query, user.Username)
	if err != nil {
		return fmt.Errorf("Error - user already existing")
	}
	return nil
}

func (repo *SQLiteUserRepo) Get(username string) (*domain.User, error) {
	query := "SELECT username FROM users WHERE username = ?"
	row := repo.db.QueryRow(query, username)
	var user domain.User
	err := row.Scan(&user.Username)
	if err != nil {
		return nil, fmt.Errorf("Error - user not found")
	}
	return &user, nil
}

func (repo *SQLiteUserRepo) Remove(user *domain.User) bool {
	query := "DELETE FROM users WHERE username = ?"
	_, err := repo.db.Exec(query, user.Username)
	if err != nil {
		fmt.Println("Error - user not found")
		return false
	}
	return true
}