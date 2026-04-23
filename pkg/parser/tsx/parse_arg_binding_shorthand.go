//ff:func feature=tsx-parser type=parser control=sequence
//ff:what parseArgBindingShorthand — `{ id }` 축약 Identifier property 에서 ArgBinding 추출

package tsx

import "encoding/json"

// parseArgBindingShorthand parses a property raw node as a shorthand
// Identifier (e.g. `{ id }` → key "id", value "id"). Returns (binding, true)
// on success; (_, false) otherwise.
func parseArgBindingShorthand(p json.RawMessage) (ArgBinding, bool) {
	var ident struct {
		Value string  `json:"value"`
		Span  astSpan `json:"span"`
	}
	if err := json.Unmarshal(p, &ident); err != nil {
		return ArgBinding{}, false
	}
	if ident.Value == "" {
		return ArgBinding{}, false
	}
	return ArgBinding{Key: ident.Value, Value: ident.Value}, true
}
