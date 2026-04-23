//ff:func feature=tsx-parser type=parser control=iteration dimension=1
//ff:what (v *visitor).descendArray — JSON array 의 모든 원소로 재귀

package tsx

import "encoding/json"

// descendArray recurses into every element of a JSON array.
func (v *visitor) descendArray(raw json.RawMessage) {
	var arr []json.RawMessage
	if err := json.Unmarshal(raw, &arr); err != nil {
		return
	}
	for _, child := range arr {
		v.walkAny(child)
	}
}
