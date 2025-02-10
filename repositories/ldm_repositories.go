// LDM (Locale Mata Management)
package repositories

import (
	// "backend/repositories"
	"fmt"
	"os"
	"path/filepath"
)

func ScanFolders(root string) ([]string, error) {
	var folders []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != root {
			// Get the relative path from root
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			// Check if it is a direct child of root
			if filepath.Dir(relPath) == "." {
				folders = append(folders, path)
			}
			// Skip subdirectories
			return filepath.SkipDir
		}
		return nil
	})
	return folders, err
}

func GetAllDataFromLocal(root string) ([]string, error) {
	folders, err := ScanFolders(root)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// slicing srings in array
	var folderCleans []string
	for _, folder := range folders {
		// KAYANYA INI GA USAH !!!
		// Convert Windows-style path to Unix-style path for correct splitting
		// unixStylePath := strings.ReplaceAll(folder, "\\", "/")

		// Extract the final element of the path
		// finalFolderName := filepath.Base(unixStylePath)
		finalFolderName := filepath.Base(folder)

		// Append the result to the folderNames slice
		folderCleans = append(folderCleans, finalFolderName)
	}

	return folderCleans, nil
}
