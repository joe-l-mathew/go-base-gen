package functions

import "os"

func FolderExists(path string) bool {
	// Use os.Stat to check if the path points to a directory
	info, err := os.Stat(path)

	// Check if there is no error and the path points to a directory
	return err == nil && info.IsDir()
}
