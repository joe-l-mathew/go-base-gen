package main

import (
	"fmt"

	"github.com/joe-l-mathew/go-base-gen.git/functions"
)

func main() {
	//ask for project name
	projectName := functions.GetUserInput("Enter the project name")
	if projectName == "" {
		fmt.Println("Project name can't be empty\nExiting...")
		return
	}
	if functions.FolderExists(projectName) {
		fmt.Println("Project with same name already exist")
		return
	}
	//ask for project github link
	githubLink := functions.GetUserInput("Enter your gihub project link")
	if githubLink == "" {
		fmt.Println("gihub project link can't be empty\nExiting...")
		return
	}
	//ask for project type
	projectType := functions.SelectProjectType()
	fmt.Println(projectType)
	//select db type
	dbtype := functions.SelectDbtype()
	fmt.Println(dbtype)

}
