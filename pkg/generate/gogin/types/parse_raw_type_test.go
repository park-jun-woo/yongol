//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestParseRawType — 토큰 정규화 (배열 마커, 파라미터, 다중 단어, 길이) 회귀

package types

import "testing"

func TestParseRawType(t *testing.T) {
	cases := []struct {
		raw       string
		wantHead  string
		wantParam string
		wantArray bool
		wantMulti bool
	}{
		{"BIGINT", "BIGINT", "", false, false},
		{"VARCHAR(255)", "VARCHAR", "255", false, false},
		{"NUMERIC(10,2)", "NUMERIC", "10,2", false, false},
		{"TEXT[]", "TEXT", "", true, false},
		{"BIGINT[]", "BIGINT", "", true, false},
		// Multi-word PG types: Head normalised to canonical alias, but
		// the informational MultiToken flag stays true.
		{"DOUBLE PRECISION", "FLOAT8", "", false, true},
		{"TIMESTAMP WITH TIME ZONE", "TIMESTAMPTZ", "", false, true},
		{"TIMESTAMP WITHOUT TIME ZONE", "TIMESTAMP", "", false, true},
		{"CHARACTER VARYING(255)", "VARCHAR", "255", false, true},
		{"TIMESTAMPTZ", "TIMESTAMPTZ", "", false, false},
		{"  varchar (10) ", "VARCHAR", "10", false, false},
	}
	for _, c := range cases {
		got := parseRawType(c.raw)
		if got.Head != c.wantHead || got.Param != c.wantParam ||
			got.IsArray != c.wantArray || got.MultiToken != c.wantMulti {
			t.Errorf("parseRawType(%q) = %+v, want head=%q param=%q array=%v multi=%v",
				c.raw, got, c.wantHead, c.wantParam, c.wantArray, c.wantMulti)
		}
	}
}
