package main

import (
	"CNAD_CloudShop/src/cli"
	"CNAD_CloudShop/src/repository"
	"CNAD_CloudShop/src/service"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/**
- REGISTER <username>
- CREATE_LISTING <username> <title> <description> <price> <category>
- DELETE_LISTING <username> <listing_id>
- GET_LISTING <username> <listing_id>
- GET_CATEGORY <username> <category>
- GET_TOP_CATEGORY <username>
*/
func parseLine(line string) []string {
    var args []string
    var current string
    var inQuotes bool

    for i := range len(line) {
        c := line[i]
        switch c {
        case ' ':
            if !inQuotes {
                if current != "" {
                    args = append(args, current)
                    current = ""
                }
                continue
            }
        case '\'', '"':
            // 切換引號狀態
            inQuotes = !inQuotes
            continue
        }
        current += string(c)
    }
    if current != "" {
        args = append(args, current)
    }
    return args
}

func main() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
		return
    }

	db, err := repository.InitDB("resources/cloudshop.db")
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}
	defer db.Close()
	
	// Initialize repositories
	startIdx, _ := strconv.ParseInt(os.Getenv("START_IDX"), 10, 64)
	userRepo := repository.NewSQLiteUserRepo(db)
	listingRepo := repository.NewSQLiteListingRepo(db, startIdx)
	categoryRepo := repository.NewSQLiteCategoryRepo(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	listingService := service.NewListingService(listingRepo, categoryRepo, userService)
	catagoryService := service.NewCategoryService(categoryRepo, userService)

	// Initialize command factory
	commandFactory := cli.NewCommandFactory(userService, listingService, catagoryService)

	// Create a scanner to read from STDIN
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// no prompt, just read the input and print the output in the expected
		fmt.Print("")
		// Read the input
		if !scanner.Scan() {
			break
		}

		// Get the line
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Parse the input
		// CREATE_LISTING user1 'Phone model 8' 'Black color, brand new' 1000 'Electronics'
		args := parseLine(line)

		// Create the command
		cmd := commandFactory.CreateCommand(args)
		if cmd == nil {
			fmt.Println("Invalid command or arguments")
			continue
		}

		// Execute the command and print the result
		cmd.Execute()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from input:", err)
	}
}