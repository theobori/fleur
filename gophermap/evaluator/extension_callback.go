package evaluator

type ExtensionCallback func(text string, e *Evaluator) (string, error)
