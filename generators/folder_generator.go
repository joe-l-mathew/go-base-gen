package generators

import (
	"fmt"
	"os"

	"github.com/joe-l-mathew/go-base-gen.git/functions"
)

func GenerateBaseFolder(folderPath string) {
	if functions.FolderExists(folderPath) {
		return
	}
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory '%s': %s\n", folderPath, err)
		os.Exit(1)
		return
	}
}

func GenerateSupportingFolders(projectName, projectType string) {
	generateFolderFiles(projectName, projectType, "middleware")
	generateFolderFiles(projectName, projectType, "routes")
	generateFolderFiles(projectName, projectType, "handlers")
	GenerateBaseFolder(projectName + "/controller")
	GenerateBaseFolder(projectName + "/config")
	GenerateBaseFolder(projectName + "/static")
}
func generateFolderFiles(projectName, projectType, folderName string) {
	GenerateBaseFolder(projectName + "/" + folderName)
	GenerateFile("templates/"+projectType+"/"+folderName+".go.tmpl", projectName+"/"+folderName+"/"+folderName+".go")
}
