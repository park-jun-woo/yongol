//ff:type feature=tsx-parser type=model
//ff:what astNode — swc AST 노드의 공통 헤더 (type + span 만 우선 디코드)

package tsx

import "encoding/json"

// astNode is a permissive view of a swc AST node. Only `type` and `span`
// are always present; all other fields are consumed on demand per node kind.
type astNode struct {
	Type string          `json:"type"`
	Span astSpan         `json:"span"`
	Raw  json.RawMessage `json:"-"`
}
