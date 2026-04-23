//ff:func feature=tsx-parser type=parser control=selection
//ff:what (v *visitor).walk — 노드 type 별 dispatch (Import / Call / KeyValueProperty)

package tsx

import "encoding/json"

// walk dispatches on node `type` and recurses into child containers. The
// walker is permissive: unknown node types are still recursed into using
// the generic value-based descent so newly-added swc node kinds never drop
// matches (e.g. inside decorators, JSX expressions).
func (v *visitor) walk(raw json.RawMessage) {
	if len(raw) == 0 || string(raw) == "null" {
		return
	}
	var head astNode
	if err := json.Unmarshal(raw, &head); err != nil {
		return
	}
	switch head.Type {
	case "ImportDeclaration":
		v.handleImport(raw)
	case "CallExpression":
		v.handleCall(raw)
	case "KeyValueProperty":
		v.handleKeyValueProperty(raw)
	}
	// Always descend — a CallExpression may host child CallExpressions (chained
	// then()s), and ImportDeclarations are leaves.
	v.descend(raw)
}
