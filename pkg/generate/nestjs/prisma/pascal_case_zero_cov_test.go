//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestPascalCase_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"user_account": "UserAccount",
		"post":         "Post",
		"a__b":         "AB",
		"":             "",
	}
	for in, want := range cases {
		if got := pascalCase(in); got != want {
			t.Errorf("pascalCase(%q)=%q want %q", in, got, want)
		}
	}
}
