//ff:type feature=ssac-parse type=model
//ff:what commentParser — type that manages comment-parsing state
package ssac

// commentParser manages the state of the comment parser.
type commentParser struct {
	sequences            []Sequence
	responseLines        []string
	inResponse           bool
	responseSuppressWarn bool
	responseStartLine    int // 1-based line number of the opening @response line
}
