package main

import (
	"fmt"
	"os"

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
	generators.GenerateBaseFolder(projectName)
	generators.GenerateSupportingFolders(projectName)
	templates := []string{
		"templates/.env.tmpl",
		"templates/Dockerfile.tmpl",
		"templates/" + projectType + "/main.go.tmpl",
	}

	for _, templatePath := range templates {
		fmt.Println(templatePath + "template path")
		// Generate the output path by replacing .tmpl with the corresponding file extension
		// outputPath := strings.TrimSuffix(templatePath, ".tmpl")
		err := generators.GenerateFile(templatePath, projectName)
		if err != nil {
			fmt.Printf("Error generating file: %s\n", err)
			os.Exit(1)
		}
	}
	functions.InitializeTheProject(githubLink, projectName)

}
