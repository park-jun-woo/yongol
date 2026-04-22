//ff:func feature=manifest type=accessor control=sequence
//ff:what LineIndex.ResponseFieldLine — operationId 의 response 필드 줄을 반환
package openapi

// ResponseFieldLine returns the line for a response body property of the given
// operationId, or 0 if unknown.
func (l *LineIndex) ResponseFieldLine(opID, field string) int {
	if l == nil {
		return 0
	}
	if m, ok := l.ResponseFields[opID]; ok {
		return m[field]
	}
	return 0
}
