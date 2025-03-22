package repository

import (
	"CNAD_CloudShop/src/domain"
	"errors"
)

var ErrCategoryNotFound = errors.New("category not found")

type UserRepo interface {
	Create(user *domain.User) error
	Get(username string) (*domain.User, error)
	Remove(user *domain.User) (bool,)
}

type ListingRepo interface {
	Create(listing *domain.Listing) (int64, error)
	Remove(username string, id int) error
	Get(id int) (*domain.Listing, error)
	GetByCategory(category string) ([]*domain.Listing, error)
	// GetTopCategories() ([]*string, error)
	GetOwner(id int) (string, error)
	// CategoryExists(category string) (bool,)
}

type CategoryRepo interface {
	Create(category *domain.Category) error
	Get(category string) (*domain.Category, error)
	Increment(category string) error
	Decrement(category string) error
	Remove(category string) error
	// GetByCategory(category string) ([]*domain.Listing, error)
	GetTopCategories() ([]*string, error)
	// GetTopCategories() ([]*string, error)
}




