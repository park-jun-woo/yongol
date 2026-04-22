//ff:func feature=manifest type=accessor control=sequence
//ff:what LineIndex.SchemaLine — components.schemas 항목의 줄 번호를 반환
package openapi

// SchemaLine returns the line for a components.schemas entry, or 0 if unknown.
func (l *LineIndex) SchemaLine(name string) int {
	if l == nil {
		return 0
	}
	return l.Schemas[name]
}
