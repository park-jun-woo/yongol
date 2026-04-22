//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what fillResponseFieldLines — 각 response 필드 제약조건에 Line 값을 채워 넣음
package openapi

// fillResponseFieldLines sets fc.Line from the LineIndex for each entry in
// fields, using the given operationId.
func fillResponseFieldLines(fields map[string]FieldConstraint, opID string, lines *LineIndex) {
	for name, fc := range fields {
		fc.Line = lines.ResponseFieldLine(opID, name)
		fields[name] = fc
	}
}
