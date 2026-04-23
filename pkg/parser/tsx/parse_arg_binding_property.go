//ff:func feature=tsx-parser type=parser control=selection
//ff:what (v *visitor).parseArgBindingProperty — apiClient 인자 ObjectExpression 의 한 property → ArgBinding

package tsx

import "encoding/json"

// parseArgBindingProperty converts a single property raw node into an
// ArgBinding. Shorthand Identifier form and KeyValueProperty form are both
// supported; every other shape returns (_, false).
func (v *visitor) parseArgBindingProperty(p json.RawMessage) (ArgBinding, bool) {
	var kv struct {
		Type string `json:"type"`
		Key  struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"key"`
		Value struct {
			Span astSpan `json:"span"`
		} `json:"value"`
	}
	if err := json.Unmarshal(p, &kv); err != nil {
		return ArgBinding{}, false
	}
	switch kv.Type {
	case "Identifier":
		return parseArgBindingShorthand(p)
	case "KeyValueProperty":
		if kv.Key.Value == "" {
			return ArgBinding{}, false
		}
		return ArgBinding{Key: kv.Key.Value, Value: v.snippet(kv.Value.Span)}, true
	}
	return ArgBinding{}, false
}
