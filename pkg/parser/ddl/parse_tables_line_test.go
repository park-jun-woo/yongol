//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what verifies that Table.Line and Table.File are populated correctly

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_FileAndLine(t *testing.T) {
	dir := t.TempDir()

	// users.sql: CREATE TABLE 이 3번째 줄
	usersSQL := `-- users table
-- for authentication
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);`
	usersPath := filepath.Join(dir, "users.sql")
	if err := os.WriteFile(usersPath, []byte(usersSQL), 0644); err != nil {
		t.Fatal(err)
	}

	// posts.sql: CREATE TABLE 이 1번째 줄
	postsSQL := `CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 0 REFERENCES users(id)
);`
	postsPath := filepath.Join(dir, "posts.sql")
	if err := os.WriteFile(postsPath, []byte(postsSQL), 0644); err != nil {
		t.Fatal(err)
	}

	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}

	byName := map[string]Table{}
	for _, tb := range tables {
		byName[tb.Name] = tb
	}

	users, ok := byName["users"]
	if !ok {
		t.Fatal("expected table 'users'")
	}
	if users.File != usersPath {
		t.Errorf("users.File = %q, want %q", users.File, usersPath)
	}
	if users.Line != 3 {
		t.Errorf("users.Line = %d, want 3", users.Line)
	}

	posts, ok := byName["posts"]
	if !ok {
		t.Fatal("expected table 'posts'")
	}
	if posts.File != postsPath {
		t.Errorf("posts.File = %q, want %q", posts.File, postsPath)
	}
	if posts.Line != 1 {
		t.Errorf("posts.Line = %d, want 1", posts.Line)
	}
}
