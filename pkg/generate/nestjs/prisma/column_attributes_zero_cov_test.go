//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestColumnAttributes_ZeroCov(t *testing.T) {
	pk := []string{"id"}
	if got := columnAttributes(ddl.Column{RawType: "BIGINT"}, "id", pk); got != "@id" {
		t.Errorf("pk attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "UUID", HasDefault: true}, "uid", pk); got != "@default(uuid())" {
		t.Errorf("uuid attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "TIMESTAMPTZ", HasDefault: true}, "ts", pk); got != "@default(now())" {
		t.Errorf("ts attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "BIGSERIAL"}, "n", pk); got != "@default(autoincrement())" {
		t.Errorf("serial attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "BIGINT", IsIdentity: true}, "n", pk); got != "@default(autoincrement())" {
		t.Errorf("identity attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "TEXT", HasDefault: true, DefaultLiteral: "x"}, "n", pk); got != `@default("x")` {
		t.Errorf("default attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "TEXT"}, "n", pk); got != "" {
		t.Errorf("no attr expected, got %q", got)
	}
}
