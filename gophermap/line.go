package gophermap

import (
	"fmt"
	"strings"
)

func isGophermapLine(line string) (bool, error) {
	values := strings.Split(line[1:], "\t")
	if len(values) != 4 {
		return false, fmt.Errorf("Invalid element number on the following Gophermap line:\n%s", line)
	}

	port := values[3]

	greaterThanZero := false
	for _, ch := range port {
		switch ch {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			greaterThanZero = true
		case '0':
			continue
		default:
			return false, fmt.Errorf("The port must be an integer with digits only.")
		}
	}

	if !greaterThanZero {
		return false, fmt.Errorf("The port must be greater than zero.")
	}

	return true, nil
}

func IsGophermapLine(line string) (bool, error) {
	if len(line) == 0 {
		return false, fmt.Errorf("The line cannot be empty.")
	}

	itemType := line[0]

	switch itemType {
	case ItemTypeTextFile, ItemTypeGopherMenu, ItemTypeCCSONameserver,
		ItemTypeErrorCode, ItemTypeBinHexEncodedFile, ItemTypeDOSBinary,
		ItemTypeUNIXUuencodedFile, ItemTypeGopherFullTextSearch,
		ItemTypeTelnet, ItemTypeBinaryFile, ItemTypeMirrorOrAlternateServer,
		ItemTypeTelnet3270, ItemTypeHTML, ItemTypeInlineText,
		ItemTypeGIFFile, ItemTypeSoundFile, ItemTypeOtherImageFile:
		// Each case above should be splitted in multiple case
		// if I'm not lazy and I decide to do custom checks
		return isGophermapLine(line)
	default:
		return false, fmt.Errorf("Invalid Gophermap item type.")
	}
}
