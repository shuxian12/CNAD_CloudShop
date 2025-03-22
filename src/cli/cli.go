package cli

import (
	"CNAD_CloudShop/src/service"
	"strconv"
)

type CommandFactory struct {
	userService    *service.UserService
	listingService *service.ListingService
	categoryService *service.CategoryService
}

func NewCommandFactory(userService *service.UserService, listingService *service.ListingService, categoryService *service.CategoryService) *CommandFactory {
	return &CommandFactory{userService, listingService, categoryService}
}

func (f *CommandFactory) CreateCommand(args []string) Command {
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "REGISTER":
		if len(args) != 2 {
			return nil
		}
		return &RegisterCommand{
			userService: f.userService,
			username:    args[1],
		}
	case "CREATE_LISTING":
		if len(args) != 6 {
			return nil
		}
		price, err := strconv.ParseInt(args[4], 10, 0)
		if err != nil {
			return nil
		}
		return &CreateListingCommand{
			listingService: f.listingService,
			username:       args[1],
			title:          trimQuotes(args[2]),
			description:    trimQuotes(args[3]),
			price:          int(price),
			category:       trimQuotes(args[5]),
		}
	case "DELETE_LISTING":
		if len(args) != 3 {
			return nil
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return nil
		}
		return &DeleteListingCommand{
			listingService: f.listingService,
			username:       args[1],
			listingID:      id,
		}
	case "GET_LISTING":
		if len(args) != 3 {
			return nil
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return nil
		}
		return &GetListingCommand{
			listService: f.listingService,
			username:       args[1],
			listingID:      id,
		}
	case "GET_CATEGORY":
		if len(args) != 3 {
			return nil
		}
		return &GetCategoryCommand{
			listingService: f.listingService,
			username:       args[1],
			category:       trimQuotes(args[2]),
		}
	case "GET_TOP_CATEGORY":
		if len(args) != 2 {
			return nil
		}
		return &GetTopCategoryCommand{
			categoryService: f.categoryService,
			username:       args[1],
		}
	}

	return nil
}

// Helper function to trim quotes from arguments
func trimQuotes(s string) string {
	if len(s) >= 2 && (s[0] == '\'' && s[len(s)-1] == '\'') || (s[0] == '"' && s[len(s)-1] == '"') {
		return s[1 : len(s)-1]
	}
	return s
}
