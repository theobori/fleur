package gophermap

import (
	"os"
	"path/filepath"
	"strings"
)

func GetDirectoryFilesText(path string, virtualPath string, domain string, port int) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	lines := make([]string, len(entries)+1)

	previousDirectoryItem := Item{
		ItemType:    ItemTypeGopherMenu,
		Description: "..",
		Selector:    filepath.Join(virtualPath, ".."),
		Domain:      domain,
		Port:        port,
	}

	lines[0] = previousDirectoryItem.String()

	for i, entry := range entries {
		entryName := entry.Name()
		entryPath := filepath.Join(path, entryName)

		var item *Item
		if entry.IsDir() {
			item, err = NewItemFromDirectoryPath(
				entryPath,
				domain,
				port,
			)
		} else {
			item, err = NewItemFromFilePath(
				entryPath,
				domain,
				port,
			)
		}

		if err != nil {
			return "", err
		}

		item.Selector = filepath.Join(virtualPath, entryName)
		item.Description = entryName

		lines[i+1] = item.String()
	}

	message := strings.Join(lines, "\n")

	return message, nil
}
