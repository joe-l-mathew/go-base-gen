package functions

import (
	"fmt"
	"os"
	"strconv"
)

func SelectDbtype() string {
	fmt.Println("Select db type")
	fmt.Println("1. Postgresql")
	fmt.Println("2. Mongodb")
	dbtype := GetUserInput("Enter the number corresponding to the database type")
	if dbtype == "" {
		return "nil"
	}
	typeNumber, err := strconv.Atoi(dbtype)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		os.Exit(1)
	}
	var dbName string
	switch typeNumber {
	case 1:
		dbName = "postgres"
	case 2:
		dbName = "mongodb"
	case 3:
		dbName = "nul"
	default:
		fmt.Println("Invalid selection. Please choose a valid project type.")
		os.Exit(1)
	}
	return dbName
}
