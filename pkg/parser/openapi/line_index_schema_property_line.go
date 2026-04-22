//ff:func feature=manifest type=accessor control=sequence
//ff:what LineIndex.SchemaPropertyLine — 스키마의 특정 property 줄 번호를 반환
package openapi

// SchemaPropertyLine returns the line for a schema property, or 0 if unknown.
func (l *LineIndex) SchemaPropertyLine(schema, prop string) int {
	if l == nil {
		return 0
	}
	if m, ok := l.SchemaProperties[schema]; ok {
		return m[prop]
	}
	return 0
}
