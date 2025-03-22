package cli

import (
	"CNAD_CloudShop/src/service"
	"fmt"
)

/**
CREATE_LISTING <username> <title> <description> <price> <category>
*/
type CreateListingCommand struct {
	listingService *service.ListingService
	username       string
	title          string
	description    string
	price          int
	category       string
}

func (clc *CreateListingCommand) Execute() {
	res, err := clc.listingService.CreateListing(clc.username, clc.title, clc.description, clc.price, clc.category)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
