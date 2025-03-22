package cli

import (
	"CNAD_CloudShop/src/service"
	"fmt"
)

/*
*
DELETE_LISTING <username> <listing_id>
*/
type DeleteListingCommand struct {
	listingService *service.ListingService
	username       string
	listingID      int
}

func (dlc *DeleteListingCommand) Execute() {
	success, err := dlc.listingService.DeleteListing(dlc.username, dlc.listingID)
	if success {
		fmt.Println("Success")
	} else {
		fmt.Println(err.Error())
	}
}
