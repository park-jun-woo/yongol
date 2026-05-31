//ff:func feature=migration type=test-helper control=sequence
//ff:what assertParseTableFKRef — parseTableFKRef 의 table/cols/consumed 결과 검증 헬퍼
package migration

import (
	"reflect"
	"testing"
)

// assertParseTableFKRef asserts the parsed table, columns, and consumed token
// count from parseTableFKRef.
func assertParseTableFKRef(t *testing.T, toks []string, wantTable string, wantCols []string, wantConsumed int) {
	t.Helper()
	gotT, gotCols, gotC := parseTableFKRef(toks)
	if gotT != wantTable {
		t.Errorf("table = %q, want %q", gotT, wantTable)
	}
	if !reflect.DeepEqual(gotCols, wantCols) {
		t.Errorf("cols = %#v, want %#v", gotCols, wantCols)
	}
	if gotC != wantConsumed {
		t.Errorf("consumed = %d, want %d", gotC, wantConsumed)
	}
}
