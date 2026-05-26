//ff:func feature=gen-typemap type=test control=iteration dimension=1
//ff:what ParseRawType — DDL 타입 문자열 정규화 검증 (head/param/array/multi-token)

package typemap

import "testing"

func TestParseRawType(t *testing.T) {
	cases := []struct {
		name      string
		raw       string
		wantHead  string
		wantParam string
		wantArray bool
		wantMulti bool
	}{
		{name: "simple", raw: "BIGINT", wantHead: "BIGINT"},
		{name: "with param", raw: "VARCHAR(255)", wantHead: "VARCHAR", wantParam: "255"},
		{name: "numeric param", raw: "NUMERIC(10,2)", wantHead: "NUMERIC", wantParam: "10,2"},
		{name: "array", raw: "TEXT[]", wantHead: "TEXT", wantArray: true},
		{name: "array with param", raw: "VARCHAR(255)[]", wantHead: "VARCHAR", wantParam: "255", wantArray: true},
		{name: "lowercase", raw: "bigint", wantHead: "BIGINT"},
		{name: "spaces", raw: "  BIGINT  ", wantHead: "BIGINT"},
		{name: "uuid", raw: "UUID", wantHead: "UUID"},
		{name: "empty", raw: "", wantHead: ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertRawType(t, c.raw, c.wantHead, c.wantParam, c.wantArray, c.wantMulti)
		})
	}
}
