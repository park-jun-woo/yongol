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
		if val, ok := matchRequiredProperty(p); ok {
			return val
		}
	}
	return false
}
