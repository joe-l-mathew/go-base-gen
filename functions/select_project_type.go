package functions

import (
	"fmt"
	"os"
	"strconv"
)

func SelectProjectType() string {
	fmt.Println("Select a project type:")
	fmt.Println("1. Gin")
	fmt.Println("2. Fiber")
	fmt.Println("3. Echo")
	projectType := GetUserInput("Enter the number corresponding to the project type")
	typeNumber, err := strconv.Atoi(projectType)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		os.Exit(1)
	}
	var projectName string
	switch typeNumber {
	case 1:
		projectName = "Gin"
	case 2:
		projectName = "Fiber"
	case 3:
		projectName = "Echo"
	default:
		fmt.Println("Invalid selection. Please choose a valid project type.")
		os.Exit(1)
	}
	return projectName
}
