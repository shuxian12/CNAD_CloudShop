package cli

import (
	"CNAD_CloudShop/src/service"
	"fmt"
)

/**
GET_TOP_CATEGORY <username>
*/
type GetTopCategoryCommand struct {
	categoryService *service.CategoryService
	username        string
}

func (gtcc *GetTopCategoryCommand) Execute() {
	res, err := gtcc.categoryService.GetTopCategory(gtcc.username)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, category := range res {
			fmt.Println(*category)
		}
	}
}
