package cli

import (
	"CNAD_CloudShop/src/service"
	"fmt"
)

/**
GET_LISTING <username> <listing_id>
*/
type GetListingCommand struct {
	listService *service.ListingService
	username    string
	listingID   int
}

func (glc *GetListingCommand) Execute() {
	listing, err := glc.listService.GetListing(glc.username, glc.listingID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v|%v|%v|%v|%v|%v\n", listing.Title, listing.Description, listing.Price, listing.CreationTime, listing.Category, listing.Username)
	}
}
