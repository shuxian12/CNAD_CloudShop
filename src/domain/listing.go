package domain

import (
	"os"
	"strings"
	"time"
)

/**
 * Listing class
 * A listing of an item put up for sale on the marketplace.
 * Only registered users should be allowed to buy or sell items.
 *
 * Listings should have the following fields:
 *  - Title
 *  - Description
 *  - Price
 *  - Username
 *  - Creation time
 *  - Category
 *  - Each listing can be associated with only 1 user and 1 category.
 */

type Listing struct {
	Title string
	Description string
	Price int
	Username string
	CreationTime string
	Category string
}

func NewListing(title string, description string, price int, username string, category string) *Listing {
	return &Listing{
		Title: title,
		Description: description,
		Price: price,
		Username: strings.ToLower(username),
		CreationTime: time.Now().Format(os.Getenv("INPUT_TIME_FORMAT")),
		Category: category,
	}
}