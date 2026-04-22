//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what fillRequestFieldLines — 각 requestBody 필드 제약조건에 Line 값을 채워 넣음
package openapi

// fillRequestFieldLines sets fc.Line from the LineIndex for each entry in
// fields, using the given operationId.
func fillRequestFieldLines(fields map[string]FieldConstraint, opID string, lines *LineIndex) {
	for name, fc := range fields {
		fc.Line = lines.RequestFieldLine(opID, name)
		fields[name] = fc
	}
}
