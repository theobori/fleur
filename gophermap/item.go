package gophermap

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/theobori/fleur/gopher"
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

func NewItemsFromBytes(b []byte) ([]*Item, error) {
	s := string(b)
	s = strings.TrimSuffix(s, gopher.CRLF+"."+gopher.CRLF)

	lines := strings.Split(s, gopher.CRLF)
	// TODO: Add GPH ??
	items, err := NewItemFromGophermapLines(lines)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func NewItemFromGophermapLines(lines []string) ([]*Item, error) {
	n := len(lines)
	items := make([]*Item, n)

	for i := range n {
		line := lines[i]
		item, err := NewItemFromGophermapLine(line)
		if err != nil {
			return nil, err
		}

		items[i] = item
	}

	return items, nil
}

func NewItemFromGophermapLine(line string) (*Item, error) {
	fields := strings.Split(line, DefaultSeparator)
	if len(fields) < 4 {
		return nil, fmt.Errorf("Invalid amount of gophermap fields in the line '%s'", line)
	}

	first := fields[0]
	lenFirst := len(first)
	if lenFirst < 1 {
		return nil, fmt.Errorf("Missing the itemtype field on the line '%s'", line)
	}
	itemType := first[0]

	var description string
	if lenFirst == 1 {
		description = ""
	} else {
		description = first[1:]
	}

	selector := fields[1]
	domain := fields[2]
	portString := fields[3]

	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, err
	}

	item, err := NewItem(itemType, description, selector, domain, port)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func NewItemFromFilePath(filePath string, domain string, port int) (*Item, error) {
	itemType, err := NewItemTypeFromFilePath(filePath)
	if err != nil {
		return nil, err
	}

	var selector string
	// We assume that if it is an HTML, CSS or JS file, it should be served via a HTTP server
	// We also assume that HTTP is available and it should automatically redirect to HTTPS
	// TODO: Maybe I should not handle it
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
