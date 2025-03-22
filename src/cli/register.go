package cli

import (
	"CNAD_CloudShop/src/service"
	"fmt"
)

type RegisterCommand struct {
	userService *service.UserService
	username    string
}

func NewRegisterCommand(userService *service.UserService, username string) *RegisterCommand {
	return &RegisterCommand{userService: userService, username: username}
}

func (rc *RegisterCommand) Execute() {
	success := rc.userService.Register(rc.username)
	if success {
		fmt.Println("Success")
	} else {
		fmt.Println("Error - user already existing")
	}
}
