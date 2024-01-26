package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	fmt.Print(prompt + ": ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim newline character from input
	return strings.TrimSpace(input)
}
