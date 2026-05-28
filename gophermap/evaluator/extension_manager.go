package evaluator

import (
	"fmt"
	"regexp"
)

// Specialized for strings, no need to be more generic for the moment
type ExtensionManager struct {
	extensions map[string]ExtensionCallback
}

func NewExtensionManager() *ExtensionManager {
	return &ExtensionManager{
		extensions: map[string]ExtensionCallback{},
	}
}

func (t *ExtensionManager) Set(pattern string, callback ExtensionCallback) {
	t.extensions[pattern] = callback
}

func (t *ExtensionManager) Delete(pattern string) bool {
	_, ok := t.extensions[pattern]
	if !ok {
		return false
	}

	delete(t.extensions, pattern)

	return true
}

func (t *ExtensionManager) Extend(text string, evaluator *Evaluator) (string, error) {
	for pattern, callback := range t.extensions {
		ok, _ := regexp.MatchString(pattern, text)
		if !ok {
			continue
		}

		ans, err := callback(text, evaluator)
		if err != nil {
			return "", err
		}

		return ans, nil
	}

	return text, fmt.Errorf("The text below could not been extended:\n%s", text)
}
