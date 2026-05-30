//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestEmitSAImports — SQLAlchemy DML(select/update/delete) import 조건부 출력

package ssac

import (
	"strings"
	"testing"
)

func TestEmitSAImports(t *testing.T) {
	cases := []struct {
		name string
		d    importData
		want string
	}{
		{"None", importData{}, ""},
		{"SelectOnly", importData{UsesSelect: true}, "from sqlalchemy import select\n"},
		{"UpdateOnly", importData{UsesUpdate: true}, "from sqlalchemy import update\n"},
		{"DeleteOnly", importData{UsesDelete: true}, "from sqlalchemy import delete\n"},
		{"All", importData{UsesSelect: true, UsesUpdate: true, UsesDelete: true}, "from sqlalchemy import select, update, delete\n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var b strings.Builder
			emitSAImports(&b, c.d)
			if b.String() != c.want {
				t.Errorf("got %q, want %q", b.String(), c.want)
			}
		})
	}
}
