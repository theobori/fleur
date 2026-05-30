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

func (t *ExtensionManager) Extend(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	for pattern, callback := range t.extensions {
		ok, _ := regexp.MatchString(pattern, ctx.Line)
		if !ok {
			continue
		}

		ans, err := callback(evaluator, ctx)
		if err != nil {
			return "", err
		}

		return ans, nil
	}

	return ctx.Line, fmt.Errorf("The text below could not been extended:\n%s", ctx.Line)
}
