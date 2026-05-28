package gophermap

import (
	"fmt"
	"os"
)

type Item struct {
	ItemType    byte
	Description string
	Selector    string
	Domain      string
	Port        int
}

func NewItem(itemType byte, description string, selector string, domain string, port int) (*Item, error) {
	if port < 0 {
		return nil, fmt.Errorf("%d is negative, the port must be positive", port)
	}

	return &Item{itemType, description, selector, domain, port}, nil
}

func (i *Item) String() string {
	return fmt.Sprintf(
		"%c%s%s%s%s%s%s%d",
		i.ItemType, i.Description,
		DefaultSeparator,
		i.Selector,
		DefaultSeparator,
		i.Domain,
		DefaultSeparator,
		i.Port,
	)
}

func NewItemFromFilePath(filePath string, domain string, port int) (*Item, error) {
	itemType, err := NewItemTypeFromFilePath(filePath)
	if err != nil {
		return nil, err
	}

	var selector string
	// We assume that if it is an HTML, CSS or JS file, it should be served via a HTTP server
	// We also assume that HTTP is available and it should automatically redirect to HTTPS
	if itemType == 'h' {
		selector = fmt.Sprintf("URL:http://%s/%s", domain, filePath)
		port = 80
	} else {
		selector = filePath
	}

	item, err := NewItem(itemType, filePath, selector, domain, port)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func NewItemFromDirectoryPath(directoryPath string, domain string, port int) (*Item, error) {
	item, err := NewItem(ItemTypeGopherMenu, directoryPath, directoryPath, domain, port)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func NewItemFromPath(path string, domain string, port int) (*Item, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return NewItemFromDirectoryPath(path, domain, port)
	}

	item, err := NewItemFromFilePath(path, domain, port)
	if err != nil {
		return nil, err
	}

	return item, nil
}
