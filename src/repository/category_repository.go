package repository

import (
	"CNAD_CloudShop/src/domain"
	"database/sql"
	"fmt"
)

type SQLiteCategoryRepo struct {
	db *sql.DB
}

func NewSQLiteCategoryRepo(db *sql.DB) CategoryRepo {
	return &SQLiteCategoryRepo{db: db}
}

// Create implements CategoryRepo.
func (repo *SQLiteCategoryRepo) Create(category *domain.Category) error {
	query := `
		INSERT INTO categories (category, count) 
		VALUES (?, ?) 
		ON CONFLICT(category) DO UPDATE SET count = count + 1
	`
	_, err := repo.db.Exec(query, category.Name, category.Count)
	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}

	return nil
}

// Get implements CategoryRepo.
func (repo *SQLiteCategoryRepo) Get(category string) (*domain.Category, error) {
	query := "SELECT category, count FROM categories WHERE category = ?"
	row := repo.db.QueryRow(query, category)

	var cat domain.Category
	err := row.Scan(&cat.Name, &cat.Count)
	if err == sql.ErrNoRows {
		return nil, ErrCategoryNotFound
	} else if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to get category: %w", err)
	}

	return &cat, nil
}

// Increment implements CategoryRepo.
func (repo *SQLiteCategoryRepo) Increment(category string) error {
	query := "UPDATE categories SET count = count + 1 WHERE category = ?"
	_, err := repo.db.Exec(query, category)
	if err != nil {
		return fmt.Errorf("failed to increment category count: %w", err)
	}
	return nil
}

// Decrement implements CategoryRepo.
func (repo *SQLiteCategoryRepo) Decrement(category string) error {
	query := "UPDATE categories SET count = count - 1 WHERE category = ?"
	_, err := repo.db.Exec(query, category)
	if err != nil {
		return fmt.Errorf("failed to decrement category count: %w", err)
	}
	return nil
}

// Remove implements CategoryRepo.
func (repo *SQLiteCategoryRepo) Remove(category string) error {
	query := "DELETE FROM categories WHERE category = ?"
	_, err := repo.db.Exec(query, category)
	if err != nil {
		return fmt.Errorf("failed to remove category: %w", err)
	}
	return nil
}

/*
GET_TOP_CATEGORY <username>
- Responses:
  - "Error - unknown user"
  - <category name> (Category having the highest total number of listings).
    This command should consider all listings in the marketplace and not just
    the listings created by the user issuing the command. Username is taken
    just for the purpose of authentication.
  - This operation is expected to be a read heavy operation as it can be used
    on the home page etc. Please ensure suitable optimization for the same.
*/
func (repo *SQLiteCategoryRepo) GetTopCategories() ([]*string, error) {
	// select category with max count
	query := `SELECT category FROM categories WHERE count = (SELECT MAX(count) FROM categories) ORDER BY category ASC;`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Cannot get top category, err: %w", err)
	}
	defer rows.Close()

	var categories []*string
	for rows.Next() {
		var category string
		rows.Scan(&category)
		categories = append(categories, &category)
	}
	if len(categories) == 0 {
		return nil, fmt.Errorf("Have no category")
	}

	return categories, nil
}
