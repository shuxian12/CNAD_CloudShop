# Cloud Shop

## Requirements
* Go: 1.24.1

## Installation
```bash
sh build.sh
```

## Usage
```bash
sh run.sh
```

## Demands
### Task
The task is to develop a solution for the below mentioned problem. You are expected to build a CLI (command line interface) application, which should take input on STDIN and post output on STDOUT (please note error messages are to be directed to STDOUT and not STDERR). There is no need for any other interface (Web, APIs etc). The key focus area will be the design of the solution and the structure of the code. Weightage will be given to clear design, extensibility of code to allow easy addition of features, modularity to ensure clear separation of concerns. Apart from the mentioned test cases, we may manually run a few additional test cases, which are closely related to the ones mentioned.

### Additional notes
Please ensure concepts of extensibility, ease of testing, abstraction and appropriate application of type safety are demonstrated.

## Architecture
```
cloudShop/
├── src/
│   ├── domain/                     # Domain models
│   │   ├── user.go
│   │   ├── listing.go
│   │   └── category.go
│   ├── repository/                 # Data access layer, persistence
│   │   ├── repository.go           # Repository interfaces
│   │   ├── sqlite_repository.go    # SQLite implementation
│   │   ├── user_repository.go
│   │   ├── listing_repository.go
│   │   └── category_repository.go
│   ├── service/                    # Business logic layer
│   │   ├── user_service.go
│   │   ├── listing_service.go
│   │   └── category_service.go
│   └── cli/                        # Command line interface
│       ├── cli.go                  # CLI runner and factory
│       ├── command.go              # Command interface
│       ├── register.go
│       ├── create_listing.go
│       ├── delete_listing.go
│       ├── get_listing.go
│       ├── get_category.go
│       └── get_top_category.go
├── resources/                      # Folder will be created at runtime
│   └── db.sqlite                   # SQLite database
├── .env                            # Environment variables
├── build.sh                        # Build script
├── run.sh                          # Run script
├── main.go                         # Entry point
├── go.mod                          # Go module definition
├── go.sum                          # Go module lock file
└── README.md                       # Project documentation
```