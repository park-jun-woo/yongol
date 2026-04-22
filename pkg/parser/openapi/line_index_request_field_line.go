//ff:func feature=manifest type=accessor control=sequence
//ff:what LineIndex.RequestFieldLine — operationId 의 requestBody 필드 줄을 반환
package openapi

// RequestFieldLine returns the line for a request body property of the given
// operationId, or 0 if unknown.
func (l *LineIndex) RequestFieldLine(opID, field string) int {
	if l == nil {
		return 0
	}
	if m, ok := l.RequestFields[opID]; ok {
		return m[field]
	}
	return 0
}
