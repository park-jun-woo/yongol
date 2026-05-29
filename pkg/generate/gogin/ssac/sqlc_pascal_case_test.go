//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what sqlcPascalCase 단위 테스트 (ID/IDS 이니셜리즘 유지, url 은 비유지)

package ssac

import "testing"

func TestSqlcPascalCase(t *testing.T) {
	cases := map[string]string{
		"id":         "ID",
		"org_id":     "OrgID",
		"created_at": "CreatedAt",
		"url":        "Url",
		"user_ids":   "UserIDS",
		"__a":        "A",
	}
	for in, want := range cases {
		if got := sqlcPascalCase(in); got != want {
			t.Errorf("sqlcPascalCase(%q) = %q, want %q", in, got, want)
		}
	}
}
