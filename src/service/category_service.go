package service

import (
	"CNAD_CloudShop/src/repository"
	"fmt"
)

type CategoryService struct {
	categoryDatabase repository.CategoryRepo
	userService      *UserService
}

func NewCategoryService(categoryRepo repository.CategoryRepo, userService *UserService) *CategoryService {
	return &CategoryService{
		categoryDatabase: categoryRepo,
		userService:      userService,
	}
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
func (s *CategoryService) GetTopCategory(username string) ([]*string, error) {
	// Check if user exists
	exists := s.userService.UserExists(username)
	if !exists {
		return nil, fmt.Errorf("Error - unknown user")
	}

	// Get top category
	categories, err := s.categoryDatabase.GetTopCategories()
	if categories == nil {
		return nil, fmt.Errorf("[ERROR] failed to get top category: %w", err)
	}

	return categories, nil
}
