//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestHasSentinelInsert(t *testing.T) {
	tests := []struct {
		name    string
		content string
		table   string
		want    bool
	}{
		{
			"sentinel present",
			"INSERT INTO users (id, name) VALUES (0, 'system');",
			"users", true,
		},
		{
			"insert but no zero sentinel",
			"INSERT INTO users (id, name) VALUES (1, 'alice');",
			"users", false,
		},
		{
			"no insert for table",
			"INSERT INTO posts (id) VALUES (0);",
			"users", false,
		},
		{
			"case insensitive keyword",
			"insert into users values (0, 'x');",
			"users", true,
		},
		{
			"empty content",
			"", "users", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSentinelInsert(tt.content, tt.table); got != tt.want {
				t.Errorf("hasSentinelInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
