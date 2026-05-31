package evaluator

import (
	"fmt"

	"github.com/theobori/fleur/gophermap"
)

func ExtendRFC1436Generic(_ *Evaluator, ctx *ExtensionContext) (string, error) {
	ok, err := gophermap.IsGophermapLine(ctx.Line)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", fmt.Errorf("This text is not RFC1436 compliant.")
	}

	return ctx.Line, nil
}

func ExtendTextFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendGopherMenuItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendCCSONameserverItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendErrorCodeItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendBinHexEncodedFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendDOSBinaryItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendUNIXUuencodedFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendGopherFullTextSearchItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendTelnetItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendBinaryFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendMirrorOrAlternateServerItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendTelnet3270Item(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendHTMLItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendInlineTextItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendGIFFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendSoundFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

func ExtendOtherImageFileItem(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	return ExtendRFC1436Generic(evaluator, ctx)
}

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
