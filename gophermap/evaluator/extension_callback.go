package evaluator

type ExtensionCallback func(e *Evaluator, ctx *ExtensionContext) (string, error)
