package repository

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"CNAD_CloudShop/src/domain"
)

type SQLiteListingRepo struct {
	db       *sql.DB
	startIdx int64
}

func NewSQLiteListingRepo(db *sql.DB, startIdx int64) ListingRepo {
	return &SQLiteListingRepo{db: db, startIdx: startIdx}
}

func (repo *SQLiteListingRepo) Create(listing *domain.Listing) (int64, error) {
	query := "INSERT INTO listings (username, title, description, price, category, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, listing.Username, listing.Title, listing.Description, listing.Price, listing.Category, listing.CreationTime)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		// fmt.Println("Error - listing not found", err)
		return -1, err
	}
	return id + repo.startIdx, nil
}

func (repo *SQLiteListingRepo) GetOwner(id int) (string, error) {
	query := "SELECT username FROM listings WHERE id = ?"
	row := repo.db.QueryRow(query, id-int(repo.startIdx))

	var owner string
	err := row.Scan(&owner)
	if err != nil {
		return "", err
	}

	return owner, nil
}

/*
DELETE_LISTING <username> <listing_id>
Responses:
  - "Success"
  - "Error - listing does not exist"
  - "Error - listing owner mismatch"
*/
func (repo *SQLiteListingRepo) Remove(username string, listingID int) error {
	query := "DELETE FROM listings WHERE id = ? AND username = ?"
	res, err := repo.db.Exec(query, listingID-int(repo.startIdx), username)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("listing does not exist, err: %w", err)
	}

	return nil
}

/*
GET_LISTING <username> <listing_id>
Responses:
  - "<title>|<description>|<price>|<created_at>|<category>|<username>"
    This command should return any listing queried according to listing_id,
    not limited to listings created by the user. Username is taken just for the
    purpose of authentication.
  - "Error - not found"
  - "Error - unknown user"
*/
func (repo *SQLiteListingRepo) Get(id int) (*domain.Listing, error) {
	query := "SELECT title, description, price, created_at, username, category FROM listings WHERE id = ?"
	row := repo.db.QueryRow(query, id-int(repo.startIdx))
	var listing domain.Listing
	var timeStr time.Time

	err := row.Scan(&listing.Title, &listing.Description, &listing.Price, &timeStr, &listing.Username, &listing.Category)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("listing not found, err: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("[ERROR] - unknown error, err: %w", err)
	}
	listing.CreationTime = timeStr.Format(os.Getenv("OUTPUT_TIME_FORMAT"))
	return &listing, nil
}

/*
*
GET_CATEGORY <username> <category>
Responses:
  - "Error - category not found"
  - "Error - unknown user"
  - "<title>|<description>|<price>|<created_at>
    <title>|<description>|<price>|<created_at>
    <title>|<description>|<price>|<created_at>
    This command should return listings of the specified category in
    descending order by create time (ie, created_at).
*/
func (repo *SQLiteListingRepo) GetByCategory(category string) ([]*domain.Listing, error) {
	query := "SELECT title, description, price, created_at, username, category FROM listings WHERE category = ? ORDER BY created_at DESC"
	rows, err := repo.db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] %w", err)
	}
	defer rows.Close()

	var listings []*domain.Listing
	for rows.Next() {
		var listing domain.Listing
		var timeStr time.Time
		rows.Scan(&listing.Title, &listing.Description, &listing.Price, &timeStr, &listing.Username, &listing.Category)
		listing.CreationTime = timeStr.Format(os.Getenv("OUTPUT_TIME_FORMAT"))
		listings = append(listings, &listing)
	}
	if len(listings) == 0 {
		return nil, fmt.Errorf("Error - category not found")
	}
	return listings, nil
}
