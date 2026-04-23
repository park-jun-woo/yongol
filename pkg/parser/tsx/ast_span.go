//ff:type feature=tsx-parser type=model
//ff:what astSpan — swc AST 노드의 byte offset 범위 (start/end, 1-based)

package tsx

// astSpan holds the 1-based byte offset range emitted by swc for every AST
// node. (line, col) resolution happens in visitor.resolve via a lineOffset index.
type astSpan struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
