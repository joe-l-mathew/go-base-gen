package functions

import (
	"fmt"
	"os"
	"os/exec"
)

func InitializeTheProject(githubLink, projectName string) {
	os.Chdir(projectName)
	cmd := exec.Command("go", "mod", "init", githubLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running 'go mod init': %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Go module initialized successfully with module name: %s\n", githubLink)
}
