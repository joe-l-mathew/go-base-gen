package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateAllBaseFiles(projectType, projectName string) {
	templates := []string{
		"templates/.env.tmpl",
		"templates/Dockerfile.tmpl",
		"templates/" + projectType + "/main.go.tmpl",
	}

	for _, templatePath := range templates {
		baseName := filepath.Base(templatePath)
		err := GenerateFile(templatePath, projectName+"/"+strings.TrimSuffix(baseName, ".tmpl"))
		if err != nil {
			fmt.Printf("Error generating file: %s\n", err)
			os.Exit(1)
		}
	}
}

func GenerateFile(templatePath, filePath string) error {
	// Read the template file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}
	// Create the output file
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	// Write the template content to the output file
	_, err = outputFile.Write(templateContent)
	if err != nil {
		return err
	}
	return nil
}
