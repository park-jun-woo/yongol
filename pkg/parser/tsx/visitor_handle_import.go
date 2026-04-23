//ff:func feature=tsx-parser type=parser control=iteration dimension=1
//ff:what (v *visitor).handleImport — 로컬 component import 만 ComponentImports 에 수집

package tsx

import "encoding/json"

// handleImport extracts ComponentImports from local paths only. npm package
// imports (react, @tanstack/*, etc) are intentionally skipped so T-1 only
// triggers on imports yongol is responsible for.
func (v *visitor) handleImport(raw json.RawMessage) {
	var d struct {
		Span       astSpan           `json:"span"`
		Specifiers []json.RawMessage `json:"specifiers"`
		Source     struct {
			Value string `json:"value"`
		} `json:"source"`
	}
	if err := json.Unmarshal(raw, &d); err != nil {
		return
	}
	if !isLocalComponentImport(d.Source.Value) {
		return
	}
	line, _ := v.resolve(d.Span.Start)
	for _, spec := range d.Specifiers {
		v.appendComponentImport(spec, d.Source.Value, line)
	}
}
