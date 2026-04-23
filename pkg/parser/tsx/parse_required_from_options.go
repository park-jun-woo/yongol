//ff:func feature=tsx-parser type=util control=iteration dimension=1
//ff:what parseRequiredFromOptions — register() 옵션 객체에서 `required: bool` 값 추출

package tsx

import "encoding/json"

// parseRequiredFromOptions extracts the boolean value of the `required`
// property from a register() options literal. Returns false on any shape
// mismatch; non-literal expressions (variables, ternaries) are treated as
// "undetermined" → false.
func parseRequiredFromOptions(arg json.RawMessage) bool {
	var w struct {
		Expression struct {
			Type       string            `json:"type"`
			Properties []json.RawMessage `json:"properties"`
		} `json:"expression"`
	}
	if err := json.Unmarshal(arg, &w); err != nil {
		return false
	}
	if w.Expression.Type != "ObjectExpression" {
		return false
	}
	for _, p := range w.Expression.Properties {
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
			continue
		}
		if kv.Type != "KeyValueProperty" {
			continue
		}
		if kv.Key.Value == "required" && kv.Value.Type == "BooleanLiteral" {
			return kv.Value.Value
		}
	}
	return false
}
