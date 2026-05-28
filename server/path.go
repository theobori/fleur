package server

import (
	"os"
	"path/filepath"
	"strings"
)

// TODO: test safePath
func SafePath(path string) string {
	st := []string{}

	splittedPath := strings.Split(path, "/")
	for _, el := range splittedPath {
		switch el {
		case ".":
			continue
		case "..":
			stLen := len(st)
			if stLen > 0 {
				st = st[:stLen-1]
			}
		default:
			st = append(st, el)
		}
	}

	return strings.Join(st, "/")
}

// TODO: test renderPersonalGopherspacePath
func RenderPersonalGopherspacePath(path string, prefixDirectoryPath string) string {
	// TODO trim prefix
	i := strings.Index(path, "~")
	if i != 0 {
		return path
	}

	// TODO handle case without / on the right
	j := strings.Index(path, "/")
	if i < 2 {
		return path
	}

	username := path[i:j]
	userHomeDirectory := filepath.Join(prefixDirectoryPath, username, "public_gopher")

	fileInfo, err := os.Stat(path)
	if err != nil {
		return path
	}

	if !fileInfo.IsDir() {
		return path
	}

	return userHomeDirectory
}
