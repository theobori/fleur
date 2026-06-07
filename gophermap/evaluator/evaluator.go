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

func (e *Evaluator) evalLine(ctx *ExtensionContext) (string, error) {
	ans, err := e.ExtensionManager.Extend(e, ctx)
	if err == nil {
		return ans, nil
	}

	if !e.options.EnableAutoInlineText {
		return "", fmt.Errorf("Unable to evaluate the line below:\n%s", ctx.Line)
	}

	item, err := gophermap.NewItem(
		gophermap.ItemTypeInlineText,
		ctx.Line,
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

func (e *Evaluator) EvalWithPathContext(source string, path string, virtualPath string) (string, error) {
	destinationLines := []string{}

	source = strings.Trim(source, "\n")
	lines := strings.SplitSeq(source, "\n")

	for sourceLine := range lines {
		ctx := ExtensionContext{
			Line:        sourceLine,
			Path:        path,
			VirtualPath: virtualPath,
		}

		destinationLine, err := e.evalLine(&ctx)
		if err != nil {
			return "", err
		}

		destinationLines = append(destinationLines, destinationLine)
	}

	ans := strings.Join(destinationLines, "\n")

	return ans, nil
}

func (e *Evaluator) EvalFile(path string, virtualPath string) (string, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return e.EvalWithPathContext(string(source), path, virtualPath)
}
