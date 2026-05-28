package evaluator

import (
	"fmt"
	"os"
	"strings"

	"github.com/theobori/fleur/gophermap"
)

type Evaluator struct {
	options          *Options
	ExtensionManager *ExtensionManager
}

func NewEvaluator(options *Options, extensionManager *ExtensionManager) *Evaluator {
	return &Evaluator{
		options:          options,
		ExtensionManager: extensionManager,
	}
}

func (e *Evaluator) Options() Options {
	return *e.options
}

func (e *Evaluator) evalLine(line string) (string, error) {
	ans, err := e.ExtensionManager.Extend(line, e)
	if err == nil {
		return ans, nil
	}

	if !e.options.EnableAutoInlineText {
		return "", fmt.Errorf("Unable to evaluate the line below:\n%s", line)
	}

	item, err := gophermap.NewItem(
		gophermap.ItemTypeInlineText,
		line,
		"/",
		e.options.Domain,
		e.options.Port,
	)
	if err != nil {
		return "", err
	}

	ans = item.String()

	return ans, nil
}

func (e *Evaluator) Eval(source string) (string, error) {
	destinationLines := []string{}
	lines := strings.Split(source, "\n")

	for _, sourceLine := range lines {
		if len(sourceLine) == 0 {
			continue
		}

		destinationLine, err := e.evalLine(sourceLine)
		if err != nil {
			return "", err
		}

		destinationLines = append(destinationLines, destinationLine)
	}

	ans := strings.Join(destinationLines, "\n")

	return ans, nil
}

func (e *Evaluator) EvalFile(path string) (string, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return e.Eval(string(source))
}
