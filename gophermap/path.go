package gophermap

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/theobori/fleur/internal/common"
)

const DefaultMaxColumns = 70

func GetDirectoryFilesAsGophermap(
	path string,
	virtualPath string,
	domain string,
	port int,
	maxColumns int,
) (string, error) {
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

		entryInfo, err := entry.Info()
		if err != nil {
			return "", err
		}

		var (
			item       *Item
			sizeString string
		)
		if entry.IsDir() {
			item, err = NewItemFromDirectoryPath(
				entryPath,
				domain,
				port,
			)
			sizeString = "-"
		} else {
			item, err = NewItemFromFilePath(
				entryPath,
				domain,
				port,
			)
			sizeString = common.GetStringFromSize(entryInfo.Size())
		}

		if err != nil {
			return "", err
		}

		entryModString := entryInfo.ModTime().String()[:16]
		rightColumn := sizeString + "   " + entryModString
		spacesAmount := max(maxColumns-len(entryName)-len(rightColumn), 1)
		spaces := strings.Repeat(" ", spacesAmount)

		item.Description = entryName + spaces + rightColumn
		item.Selector = filepath.Join(virtualPath, entryName)

		lines[i+1] = item.String()
	}

	message := strings.Join(lines, "\n")

	return message, nil
}
