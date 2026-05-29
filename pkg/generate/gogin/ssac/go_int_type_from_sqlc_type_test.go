//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what goIntTypeFromSqlcType 단위 테스트

package ssac

import "testing"

func TestGoIntTypeFromSqlcType(t *testing.T) {
	cases := map[string]string{
		"pgtype.Int8": "int64",
		"pgtype.Int4": "int32",
		"pgtype.Int2": "int16",
		"int64":       "int64",
		"int32":       "int32",
		"pgtype.Text": "",
		"":            "",
	}
	for in, want := range cases {
		if got := goIntTypeFromSqlcType(in); got != want {
			t.Errorf("goIntTypeFromSqlcType(%q) = %q, want %q", in, got, want)
		}
	}
}
