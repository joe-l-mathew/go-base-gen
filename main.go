package main

import (
	"fmt"

	"github.com/joe-l-mathew/go-base-gen.git/functions"
	"github.com/joe-l-mathew/go-base-gen.git/generators"
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
	githubLink := functions.GetUserInput("Enter your gihub project link or mod name")
	if githubLink == "" {
		fmt.Println("gihub project link can't be empty\nExiting...")
		return
	}
	//ask for project type
	projectType := functions.SelectProjectType()
	//select db type
	// dbtype := functions.SelectDbtype()
	//create the base folder with project name
	generators.GenerateBaseFolder(projectName)
	// create all the internal folders
	generators.GenerateSupportingFolders(projectName, projectType)
	//running go mod init
	defer functions.InstallRouters(projectType)
	defer functions.InitializeTheProject(githubLink, projectName)
	generators.GenerateAllBaseFiles(projectType, projectName)

}
