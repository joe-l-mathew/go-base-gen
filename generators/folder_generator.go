package generators

import (
	"fmt"
	"os"

	"github.com/joe-l-mathew/go-base-gen.git/functions"
)

func GenerateBaseFolder(projectName string) {
	if functions.FolderExists(projectName) {
		return
	}
	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory '%s': %s\n", projectName, err)
		os.Exit(1)
		return
	}
}

func GenerateSupportingFolders(projectName string) {
	GenerateBaseFolder(projectName + "/middleware")
	GenerateBaseFolder(projectName + "/routes")
	GenerateBaseFolder(projectName + "/controller")
	GenerateBaseFolder(projectName + "/config")
	GenerateBaseFolder(projectName + "/handler")
	GenerateBaseFolder(projectName + "/static")
}
