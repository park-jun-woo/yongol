//ff:func feature=tsx-parser type=parser control=selection
//ff:what (v *visitor).walkAny — JSON 조각 형태(object/array/scalar) 에 맞는 순회 경로 선택

package tsx

import "encoding/json"

// walkAny routes into walk for object-shaped AST nodes and into descend
// for arrays. Scalar leaves are skipped.
func (v *visitor) walkAny(raw json.RawMessage) {
	if len(raw) == 0 {
		return
	}
	switch raw[0] {
	case '{':
		v.walk(raw)
	case '[':
		v.descend(raw)
	}
}
