package evaluator

import (
	"fmt"
	"regexp"
)

const ExtensionMinimumMaxWeight = 10

// Specialized for strings, no need to be more generic for the moment
type ExtensionManager struct {
	extensions []map[string]ExtensionCallback
	maxWeight  int
}

func NewExtensionManagerWithMaxWeight(maxWeight int) (*ExtensionManager, error) {
	if maxWeight < ExtensionMinimumMaxWeight {
		return nil, fmt.Errorf("The minimum weight allowed is %d", ExtensionMinimumMaxWeight)
	}

	extensions := []map[string]ExtensionCallback{}
	for range maxWeight + 1 {
		extensions = append(extensions, map[string]ExtensionCallback{})
	}
	return &ExtensionManager{
		extensions: extensions,
		maxWeight:  maxWeight,
	}, nil
}

func NewExtensionManager() *ExtensionManager {
	extensionManager, _ := NewExtensionManagerWithMaxWeight(ExtensionMinimumMaxWeight)

	return extensionManager
}

func (e *ExtensionManager) getIndexFromWeight(weight int) (int, error) {
	if weight < 0 || weight > e.maxWeight {
		return -1, fmt.Errorf("Weight must be between %d and %d", 0, e.maxWeight)
	}

	i := e.maxWeight - weight

	return i, nil
}

func (e *ExtensionManager) SetWithWeight(weight int, pattern string, callback ExtensionCallback) error {
	i, err := e.getIndexFromWeight(weight)
	if err != nil {
		return err
	}

	e.extensions[i][pattern] = callback

	return nil
}

func (e *ExtensionManager) Set(pattern string, callback ExtensionCallback) {
	e.SetWithWeight(e.maxWeight, pattern, callback)
}

func (e *ExtensionManager) DeleteWithWeight(weight int, pattern string, callback ExtensionCallback) (bool, error) {
	i, err := e.getIndexFromWeight(weight)
	if err != nil {
		return false, err
	}

	_, ok := e.extensions[i][pattern]
	if !ok {
		return false, nil
	}

	delete(e.extensions[i], pattern)

	return true, nil
}

func (e *ExtensionManager) Delete(pattern string, callback ExtensionCallback) bool {
	ok, _ := e.DeleteWithWeight(e.maxWeight, pattern, callback)

	return ok
}

func (e *ExtensionManager) Extend(evaluator *Evaluator, ctx *ExtensionContext) (string, error) {
	for _, extensions := range e.extensions {
		for pattern, callback := range extensions {
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
	}

	return ctx.Line, fmt.Errorf("The text below could not been extended:\n%s", ctx.Line)
}
