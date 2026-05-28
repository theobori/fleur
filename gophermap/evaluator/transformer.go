package evaluator

import (
	"fmt"

	"github.com/theobori/fleur/gophermap"
)

func RFC1436ItemsExtensionManager() *ExtensionManager {
	t := NewExtensionManager()
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeTextFile), ExtendTextFileItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeGopherMenu), ExtendGopherMenuItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeCCSONameserver), ExtendCCSONameserverItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeErrorCode), ExtendErrorCodeItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeBinHexEncodedFile), ExtendBinHexEncodedFileItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeDOSBinary), ExtendDOSBinaryItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeUNIXUuencodedFile), ExtendUNIXUuencodedFileItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeGopherFullTextSearch), ExtendGopherFullTextSearchItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeTelnet), ExtendTelnetItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeBinaryFile), ExtendBinaryFileItem)
	t.Set(fmt.Sprintf("^\\%c.*", gophermap.ItemTypeMirrorOrAlternateServer), ExtendMirrorOrAlternateServerItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeTelnet3270), ExtendTelnet3270Item)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeHTML), ExtendHTMLItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeInlineText), ExtendInlineTextItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeGIFFile), ExtendGIFFileItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeSoundFile), ExtendSoundFileItem)
	t.Set(fmt.Sprintf("^%c.*", gophermap.ItemTypeOtherImageFile), ExtendOtherImageFileItem)

	return t
}
