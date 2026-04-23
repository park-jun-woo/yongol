//ff:func feature=tsx-parser type=parser control=selection
//ff:what (v *visitor).descend — JSON object/array 의 모든 자식 값으로 재귀 (type/span 필드는 스킵)

package tsx

import "encoding/json"

// descend recurses into every child value of an arbitrary JSON object or
// array, invoking walk on every encountered object. This is simpler than
// enumerating swc's ~120 node shapes and the overhead is negligible for
// per-page ASTs (typically a few hundred KB).
func (v *visitor) descend(raw json.RawMessage) {
	// Fast reject: only objects / arrays contain children.
	switch raw[0] {
	case '{':
		v.descendObject(raw)
	case '[':
		v.descendArray(raw)
	}
}
