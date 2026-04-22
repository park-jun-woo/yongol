//ff:func feature=manifest type=accessor control=sequence
//ff:what LineIndex.OperationLine — operationId 의 등장 줄을 반환
package openapi

// OperationLine returns the line for the given operationId, or 0 if unknown.
func (l *LineIndex) OperationLine(opID string) int {
	if l == nil {
		return 0
	}
	return l.Operations[opID]
}
