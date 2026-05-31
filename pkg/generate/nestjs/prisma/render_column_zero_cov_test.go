//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderColumn_ZeroCov(t *testing.T) {
	var b strings.Builder
	// PK column (not optional, @id)
	renderColumn(&b, ddl.Column{RawType: "BIGINT", NotNull: true}, "id", []string{"id"})
	// nullable column, no index → optional "?"
	renderColumn(&b, ddl.Column{RawType: "TEXT"}, "bio", []string{"id"})
	// unique column with attrs
	idxs := []ddl.Index{{Columns: []string{"email"}, IsUnique: true}}
	renderColumn(&b, ddl.Column{RawType: "TEXT", NotNull: true}, "email", []string{"id"}, idxs)
	// unique column without other attrs
	idxs2 := []ddl.Index{{Columns: []string{"slug"}, IsUnique: true}}
	renderColumn(&b, ddl.Column{RawType: "TEXT", NotNull: true}, "slug", []string{"id"}, idxs2)
	out := b.String()
	if !strings.Contains(out, "@id") {
		t.Error("expected @id")
	}
	if !strings.Contains(out, "String?") {
		t.Error("expected optional String?")
	}
	if !strings.Contains(out, "@unique") {
		t.Error("expected @unique")
	}
}
