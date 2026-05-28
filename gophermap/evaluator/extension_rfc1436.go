package evaluator

import (
	"fmt"

	"github.com/theobori/fleur/gophermap"
)

func extendRFC1436Generic(text string, evaluator *Evaluator) (string, error) {
	ok, err := gophermap.IsGophermapLine(text)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", fmt.Errorf("This text is not RFC1436 compliant.")
	}

	return text, nil
}

func ExtendTextFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendGopherMenuItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendCCSONameserverItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendErrorCodeItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendBinHexEncodedFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendDOSBinaryItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendUNIXUuencodedFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendGopherFullTextSearchItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendTelnetItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendBinaryFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendMirrorOrAlternateServerItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendTelnet3270Item(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendHTMLItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendInlineTextItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendGIFFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendSoundFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}

func ExtendOtherImageFileItem(text string, evaluator *Evaluator) (string, error) {
	return extendRFC1436Generic(text, evaluator)
}
