//ff:func feature=tsx-parser type=parser control=sequence
//ff:what (v *visitor).appendComponentImport — 단일 import specifier → ComponentImport 으로 변환해 page 에 추가

package tsx

import "encoding/json"

// appendComponentImport parses a single import specifier raw node and, when
// it resolves to a named local/imported identifier, appends the resulting
// ComponentImport to v.page.Imports.
func (v *visitor) appendComponentImport(spec json.RawMessage, source string, line int) {
	var s struct {
		Type  string `json:"type"`
		Local struct {
			Value string `json:"value"`
		} `json:"local"`
		Imported *struct {
			Value string `json:"value"`
		} `json:"imported"`
	}
	if err := json.Unmarshal(spec, &s); err != nil {
		return
	}
	name := s.Local.Value
	if s.Imported != nil && s.Imported.Value != "" {
		name = s.Imported.Value
	}
	if name == "" {
		return
	}
	v.page.Imports = append(v.page.Imports, ComponentImport{
		Name: name,
		Path: source,
		Line: line,
	})
}
