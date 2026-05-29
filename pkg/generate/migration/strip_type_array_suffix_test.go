//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestStripTypeArraySuffix — [] 접미 반복 제거 + array 여부 반환
package migration

import "testing"

func TestStripTypeArraySuffix(t *testing.T) {
	cases := []struct {
		in        string
		wantArray bool
		wantTrim  string
	}{
		{"INTEGER", false, "INTEGER"},
		{"INTEGER[]", true, "INTEGER"},
		{"INTEGER[][]", true, "INTEGER"},
		{"TEXT []", true, "TEXT"},
	}
	for _, c := range cases {
		array, trim := stripTypeArraySuffix(c.in)
		if array != c.wantArray || trim != c.wantTrim {
			t.Errorf("stripTypeArraySuffix(%q) = (%v,%q), want (%v,%q)", c.in, array, trim, c.wantArray, c.wantTrim)
		}
	}
}
