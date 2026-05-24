//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what TestExtractSelectColumns — SELECT 절에서 * 여부 및 컬럼 목록 추출 검증

package sqlc

import (
	"testing"
)

func TestExtractSelectColumns(t *testing.T) {
	cases := []struct {
		name     string
		body     string
		wantStar bool
		wantCols []string
	}{
		{
			name:     "SELECT *",
			body:     "SELECT * FROM users WHERE id = @id",
			wantStar: true,
			wantCols: nil,
		},
		{
			name:     "SELECT * multiline",
			body:     "SELECT *\nFROM users\nWHERE id = @id",
			wantStar: true,
			wantCols: nil,
		},
		{
			name:     "table-qualified SELECT *",
			body:     "SELECT u.* FROM users u WHERE u.id = @id",
			wantStar: true,
			wantCols: nil,
		},
		{
			name:     "specific columns",
			body:     "SELECT id, email, name FROM users WHERE id = @id",
			wantStar: false,
			wantCols: []string{"id", "email", "name"},
		},
		{
			name:     "specific columns multiline",
			body:     "SELECT id, email, name\nFROM users\nWHERE id = @id",
			wantStar: false,
			wantCols: []string{"id", "email", "name"},
		},
		{
			name:     "table-qualified columns",
			body:     "SELECT u.id, u.token_hash, u.expires_at FROM refresh_tokens u WHERE u.token_hash = @token_hash",
			wantStar: false,
			wantCols: []string{"id", "token_hash", "expires_at"},
		},
		{
			name:     "aliased column",
			body:     "SELECT id, count(*) AS total FROM users GROUP BY id",
			wantStar: false,
			wantCols: []string{"id", "total"},
		},
		{
			name:     "INSERT (no SELECT)",
			body:     "INSERT INTO users (email) VALUES (@email) RETURNING *",
			wantStar: false,
			wantCols: nil,
		},
		{
			name:     "UPDATE (no SELECT)",
			body:     "UPDATE users SET name = @name WHERE id = @id RETURNING *",
			wantStar: false,
			wantCols: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			star, cols := extractSelectColumns(tc.body)
			assertSelectColumns(t, star, cols, tc.wantStar, tc.wantCols)
		})
	}
}
