//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what parseCreateTableName — CREATE TABLE 헤더에서 table name 추출 (IF NOT EXISTS/인용/소문자화)

package ddl

import "testing"

func TestParseCreateTableName(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"CREATE TABLE users (", "users"},
		{"CREATE TABLE IF NOT EXISTS Accounts (", "accounts"},
		{`CREATE TABLE "MixedCase" (`, "mixedcase"},
		{"create table widgets(", "widgets"},
		{"CREATE TABLE lonely;", "lonely"},
		{"ALTER TABLE users", ""},
		{"CREATE TABLE trailing", "trailing"},
	}
	for _, c := range cases {
		if got := parseCreateTableName(c.in); got != c.want {
			t.Errorf("parseCreateTableName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
