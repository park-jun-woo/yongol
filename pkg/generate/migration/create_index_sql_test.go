//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCreateIndex_SQL — UNIQUE/USING/WHERE 절 조합 렌더
package migration

import (
	"testing"
)

func TestCreateIndex_SQL(t *testing.T) {
	cases := []struct {
		name string
		op   CreateIndex
		want string
	}{
		{
			"plain",
			CreateIndex{Table: "users", Index: &Index{Name: "idx_email", Columns: []string{"email"}}},
			"CREATE INDEX idx_email ON users (email);",
		},
		{
			"unique multi-col",
			CreateIndex{Table: "users", Index: &Index{Name: "u", Columns: []string{"a", "b"}, Unique: true}},
			"CREATE UNIQUE INDEX u ON users (a, b);",
		},
		{
			"using gin with where",
			CreateIndex{Table: "docs", Index: &Index{Name: "g", Columns: []string{"body"}, Method: "gin", Where: "active"}},
			"CREATE INDEX g ON docs USING gin (body) WHERE active;",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := c.op.SQL(); got != c.want {
				t.Errorf("SQL() = %q, want %q", got, c.want)
			}
		})
	}
}
