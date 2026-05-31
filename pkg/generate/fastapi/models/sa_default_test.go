//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestSaDefault(t *testing.T) {
	cases := []struct {
		col  ddl.Column
		want string
	}{
		{ddl.Column{RawType: "UUID", HasDefault: true}, "default=uuid.uuid4"},
		{ddl.Column{RawType: "TIMESTAMPTZ", HasDefault: true}, `server_default="now()"`},
		{ddl.Column{RawType: "DATE", HasDefault: true}, `server_default="now()"`},
		{ddl.Column{RawType: "SERIAL"}, ""},
		{ddl.Column{RawType: "BIGSERIAL"}, ""},
		{ddl.Column{RawType: "TEXT", DefaultLiteral: ""}, ""},
		{ddl.Column{RawType: "BOOLEAN", HasDefault: true, DefaultLiteral: "true"}, "default=True"},
		{ddl.Column{RawType: "BOOLEAN", HasDefault: true, DefaultLiteral: "FALSE"}, "default=False"},
		{ddl.Column{RawType: "INTEGER", HasDefault: true, DefaultLiteral: "42"}, "default=42"},
		{ddl.Column{RawType: "TEXT", HasDefault: true, DefaultLiteral: "hello"}, `default="hello"`},
	}
	for _, c := range cases {
		if got := saDefault(c.col); got != c.want {
			t.Errorf("saDefault(%+v) = %q, want %q", c.col, got, c.want)
		}
	}
}
