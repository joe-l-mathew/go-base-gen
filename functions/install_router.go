package functions

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallRouters(framework string) error {
	var packageName string
	switch framework {
	case "Gin":
		packageName = "github.com/gin-gonic/gin"
	case "Fiber":
		packageName = "github.com/gofiber/fiber/v2"
	case "Echo":
		packageName = "github.com/labstack/echo/v4"
	default:
		fmt.Println("Invalid framework choice. Supported frameworks: gin, fiber, echo.")
		os.Exit(1)
	}
	cmd := exec.Command("go", "get", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
