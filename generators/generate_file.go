package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFile(templatePath, projectName string) error {
	baseName := filepath.Base(templatePath)

	// Read the template file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Create the output file
	outputFile, err := os.Create(projectName + "/" + strings.TrimSuffix(baseName, ".tmpl"))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write the template content to the output file
	_, err = outputFile.Write(templateContent)
	if err != nil {
		return err
	}

	fmt.Printf("File '%s' generated successfully.\n", projectName)
	return nil
}
