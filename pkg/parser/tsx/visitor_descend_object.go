//ff:func feature=tsx-parser type=parser control=iteration dimension=1
//ff:what (v *visitor).descendObject — JSON object 의 자식 값으로 재귀 (type/span 키는 스킵)

package tsx

import "encoding/json"

// descendObject recurses into every non-metadata child of a JSON object.
// `type` and `span` keys carry node metadata handled by walk, so they are
// skipped here.
func (v *visitor) descendObject(raw json.RawMessage) {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil {
		return
	}
	for k, child := range obj {
		if k == "type" || k == "span" {
			continue
		}
		v.walkAny(child)
	}
}
