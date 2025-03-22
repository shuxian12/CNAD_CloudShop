package cli

import (
	"CNAD_CloudShop/src/service"
	"fmt"
)

/**
GET_CATEGORY <username> <category>

# GET_CATEGORY user1 'Sports'
T-shirt|White color|20|2019-02-22 12:34:58|Sports|user2
Black shoes|Training shoes|100|2019-02-22 12:34:57|Sports|user1
*/
type GetCategoryCommand struct {
	listingService *service.ListingService
	username        string
	category        string
}

func (gcc *GetCategoryCommand) Execute() {
	res, err := gcc.listingService.GetByCategory(gcc.username, gcc.category)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, listing := range res {
			fmt.Printf("%v|%v|%v|%v|%v|%v\n",
				listing.Title,
				listing.Description,
				listing.Price,
				listing.CreationTime,
				listing.Category,
				listing.Username)
		}
	}
}
