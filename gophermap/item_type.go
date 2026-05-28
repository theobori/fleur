package gophermap

import (
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

// See gopher://baud.baby/0/phlog/fs20181102.txt
const (
	ItemTypeTextFile                byte = '0'
	ItemTypeGopherMenu              byte = '1' // Submenu or directory
	ItemTypeCCSONameserver          byte = '2'
	ItemTypeErrorCode               byte = '3'
	ItemTypeBinHexEncodedFile       byte = '4'
	ItemTypeDOSBinary               byte = '5'
	ItemTypeUNIXUuencodedFile       byte = '6'
	ItemTypeGopherFullTextSearch    byte = '7'
	ItemTypeTelnet                  byte = '8'
	ItemTypeBinaryFile              byte = '9'
	ItemTypeMirrorOrAlternateServer byte = '+'
	ItemTypeTelnet3270              byte = 'T'
	ItemTypeHTML                    byte = 'h'
	ItemTypeInlineText              byte = 'i'
	ItemTypeGIFFile                 byte = 'g'
	ItemTypeSoundFile               byte = 's'
	ItemTypeOtherImageFile          byte = 'I'
)

func NewItemTypeFromFileExtension(extension string) byte {
	switch extension {
	case ".gophermap":
		return ItemTypeGopherMenu
	case ".txt", ".md":
		return ItemTypeTextFile
	case ".gif":
		return ItemTypeGIFFile
	case ".html", ".css", ".js":
		return ItemTypeHTML
	case ".bin", ".img", ".iso":
		return ItemTypeBinaryFile
	case ".hqx":
		return ItemTypeBinHexEncodedFile
	case ".ph":
		return ItemTypeCCSONameserver
	case ".exe", ".com", ".dll", ".sys", ".pagefind":
		return ItemTypeDOSBinary
	case ".uu", ".uue":
		return ItemTypeUNIXUuencodedFile
	// See https://en.wikipedia.org/wiki/Audio_file_format
	case ".3gp", ".aa", ".aac", ".aax", ".act", ".aiff", ".alac",
		".amr", ".ape", ".au", ".awb", ".dss", ".dvf", ".flac",
		".gsm", ".iklax", ".ivs", ".m4a", ".m4b", ".m4p", ".mmf",
		".movpkg", ".mp1", ".mp2", ".mp3", ".mpc", ".msv",
		".nmf", ".ogg", ".opus", ".ra", ".raw", ".rf64", ".sln",
		".tta", ".voc", ".vox", ".wav", ".wma", ".wv", ".webm",
		".8svx", ".cda":
		return ItemTypeSoundFile
	// See https://developer.mozilla.org/fr/docs/Web/Media/Guides/Formats/Image_types
	case ".apng", ".avif", ".jpg", ".jpeg", ".jfif", ".pjpeg",
		".pjp", ".png", ".svg", ".webp", ".bmp",
		".ico", ".cur", ".tif", ".tiff":
		return ItemTypeOtherImageFile
	default:
		// Here I'm assuming that >90% of the users
		// will serve human readable files.
		//
		// So it should be read as a text file to show its raw content.
		return ItemTypeTextFile
	}
}

func NewItemTypeFromFilePath(filePath string) (byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return byte(0), err
	}

	fileName := filepath.Base(filePath)
	if fileName == DefaultIndexFileName {
		return ItemTypeGopherMenu, nil
	}

	head := make([]byte, 261)
	file.Read(head)

	kind, err := filetype.Match(head)
	if err != nil {
		return byte(0), err
	}

	extension := kind.Extension
	if extension == "unknown" {
		extension = filepath.Ext(filePath)
	}

	itemType := NewItemTypeFromFileExtension(extension)

	return itemType, nil
}
