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
			continue
		}
		// Shorthand: { id } → Type="Identifier" with same Value as key.
		if kv.Type == "Identifier" {
			var ident struct {
				Value string  `json:"value"`
				Span  astSpan `json:"span"`
			}
			if err := json.Unmarshal(p, &ident); err == nil && ident.Value != "" {
				out = append(out, ArgBinding{Key: ident.Value, Value: ident.Value})
			}
			continue
		}
		if kv.Type != "KeyValueProperty" || kv.Key.Value == "" {
			continue
		}
		out = append(out, ArgBinding{
			Key:   kv.Key.Value,
			Value: v.snippet(kv.Value.Span),
		})
	}
	return out
}
