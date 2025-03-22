package service

import (
	"CNAD_CloudShop/src/domain"
	"CNAD_CloudShop/src/repository"
	"fmt"
)

type ListingService struct {
	listingDatabase  repository.ListingRepo
	categoryDatabase repository.CategoryRepo
	userService      *UserService
}

func NewListingService(listingDatabase repository.ListingRepo, categoryDatabase repository.CategoryRepo, userService *UserService) *ListingService {
	return &ListingService{listingDatabase: listingDatabase, categoryDatabase: categoryDatabase, userService: userService}
}

/*
- CREATE_LISTING <username> <title> <description> <price> <category>
Responses:
	- "<listing id>"
	- "Error - unknown user"
*/
func (s *ListingService) CreateListing(username string, title string, description string, price int, category string) (int64, error) {
	// Check if user exists
	if !s.userService.UserExists(username) {
		return -1, fmt.Errorf("Error - unknown user")
	}

	// Create listing
	listing := domain.NewListing(title, description, price, username, category)
	id, err := s.listingDatabase.Create(listing)
	if err != nil {
		return -1, fmt.Errorf("[ERROR] %v", err)
	}

	// Update category count
	err = s.categoryDatabase.Create(&domain.Category{Name: category, Count: 1})
	if err != nil {
		return -1, fmt.Errorf("[ERROR] %v", err)
	}

	return id, nil
}

/*
- DELETE_LISTING <username> <listing_id>
Responses:
  - "Success"
  - "Error - listing does not exist"
  - "Error - listing owner mismatch"
*/
func (s *ListingService) DeleteListing(username string, listingID int) (bool, error) {
	// Check if listing exists and belongs to the user
	listing, err := s.listingDatabase.Get(listingID)
	if err != nil {
		return false, fmt.Errorf("Error - listing does not exist")
	} else if listing.Username != username {
		return false, fmt.Errorf("Error - listing owner mismatch")
	}

	// Get category
	category, err := s.categoryDatabase.Get(listing.Category)
	if err != nil {
		return false, err
	}
	if category.Count == 1 {
		s.categoryDatabase.Remove(listing.Category)
	} else {
		s.categoryDatabase.Decrement(listing.Category)
	}

	// Delete listing
	if err := s.listingDatabase.Remove(username, listingID); err != nil {
		// fmt.Println("[ERROR] %v", err)
		return false, fmt.Errorf("[ERROR] %v", err)
	}
	return true, nil
}

/**
GET_LISTING <username> <listing_id>
- Responses:
  - "<title>|<description>|<price>|<created_at>|<category>|<username>"
    This command should return any listing queried according to listing_id,
    not limited to listings created by the user. Username is taken just for the
    purpose of authentication.
  - "Error - not found"
  - "Error - unknown user"
*/
func (s *ListingService) GetListing(username string, listingID int) (*domain.Listing, error) {
	// Check if user exists
	exist := s.userService.UserExists(username)
	if !exist {
		return nil, fmt.Errorf("Error - unknown user")
	}

	// Get listing
	listing, err := s.listingDatabase.Get(listingID)
	if err != nil {
		return nil, fmt.Errorf("Error - not found")
	}

	return listing, nil
}

/*
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
func (s *ListingService) GetByCategory(username string, category string) ([]*domain.Listing, error) {
	// Check if user exists
	exists := s.userService.UserExists(username)
	if !exists {
		return nil, fmt.Errorf("Error - unknown user")
	}

	// Get listings by category
	_, err := s.categoryDatabase.Get(category)
	if err == repository.ErrCategoryNotFound {
		return nil, fmt.Errorf("Error - category not found")
	}

	listings, err := s.listingDatabase.GetByCategory(category)
	if err != nil {
		return nil, fmt.Errorf("Error - category not found")
	}
	return listings, nil
}
