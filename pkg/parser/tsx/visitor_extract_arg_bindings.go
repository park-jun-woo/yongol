//ff:func feature=tsx-parser type=parser control=iteration dimension=1
//ff:what (v *visitor).extractArgBindings — apiClient 호출 첫 인자의 top-level ObjectExpression 키 추출

package tsx

import "encoding/json"

// extractArgBindings pulls keys out of the first argument of an apiClient
// call. Only the top-level ObjectExpression's keys are captured — nested
// objects and spreads are skipped (XOT-2 compares flat parameter names only).
func (v *visitor) extractArgBindings(arg json.RawMessage) []ArgBinding {
	var w struct {
		Expression struct {
			Type       string            `json:"type"`
			Properties []json.RawMessage `json:"properties"`
		} `json:"expression"`
	}
	if err := json.Unmarshal(arg, &w); err != nil {
		return nil
	}
	if w.Expression.Type != "ObjectExpression" {
		return nil
	}
	out := make([]ArgBinding, 0, len(w.Expression.Properties))
	for _, p := range w.Expression.Properties {
		if b, ok := v.parseArgBindingProperty(p); ok {
			out = append(out, b)
		}
	}
	return out
}
