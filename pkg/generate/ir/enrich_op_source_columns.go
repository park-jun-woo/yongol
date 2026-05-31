//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what enrichOpSourceColumns -- 단일 Op 의 모든 FieldArg.Field 에서 SourceColumn(snake_case) 세팅

package ir

// enrichOpSourceColumns sets SourceColumn on every FieldArg of op that has a
// Field accessor. DDL-independent (Pass 1).
func enrichOpSourceColumns(op *Op) {
	for _, args := range collectFieldArgSlices(op) {
		setSourceColumns(*args)
	}
}
