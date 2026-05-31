//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what assertLookupDDLColumn — lookupDDLColumn 결과(nil 여부 / RawType) 검증 헬퍼
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// assertLookupDDLColumn asserts lookupDDLColumn's nil-ness and RawType.
func assertLookupDDLColumn(t *testing.T, tables []ddl.Table, model, column string, wantNil bool, wantRaw string) {
	t.Helper()
	got := lookupDDLColumn(tables, model, column)
	if wantNil {
		if got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
		return
	}
	if got == nil {
		t.Fatalf("expected column, got nil")
	}
	if got.RawType != wantRaw {
		t.Errorf("RawType = %q, want %q", got.RawType, wantRaw)
	}
}
