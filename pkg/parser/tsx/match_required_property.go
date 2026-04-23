//ff:func feature=tsx-parser type=util control=sequence
//ff:what matchRequiredProperty — register() 옵션 ObjectExpression 의 단일 property 를 required 키 여부로 판정

package tsx

import "encoding/json"

// matchRequiredProperty inspects a single property of a register() options
// ObjectExpression. Returns (value, true) when the property is
// `required: <BooleanLiteral>`, otherwise (false, false).
func matchRequiredProperty(p json.RawMessage) (bool, bool) {
	var kv struct {
		Type string `json:"type"`
		Key  struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"key"`
		Value struct {
			Type  string `json:"type"`
			Value bool   `json:"value"`
		} `json:"value"`
	}
	if err := json.Unmarshal(p, &kv); err != nil {
		return false, false
	}
	if kv.Type != "KeyValueProperty" {
		return false, false
	}
	if kv.Key.Value != "required" || kv.Value.Type != "BooleanLiteral" {
		return false, false
	}
	return kv.Value.Value, true
}
