//ff:func feature=manifest type=accessor control=sequence
//ff:what LineIndex.PathLine — paths: 아래 path key 줄을 반환
package openapi

// PathLine returns the line for a path key under paths:, or 0 if unknown.
func (l *LineIndex) PathLine(path string) int {
	if l == nil {
		return 0
	}
	return l.Paths[path]
}
