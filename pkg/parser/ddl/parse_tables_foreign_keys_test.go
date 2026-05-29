//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — inline REFERENCES + CONSTRAINT FOREIGN KEY 둘 다 수집

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_ForeignKeys(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY
);

CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    editor_id BIGINT,
    CONSTRAINT fk_editor FOREIGN KEY (editor_id) REFERENCES users(id)
);`
	if err := os.WriteFile(filepath.Join(dir, "schema.sql"), []byte(sql), 0o644); err != nil {
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
	posts, ok := byName["posts"]
	if !ok {
		t.Fatal("posts missing")
	}
	if len(posts.ForeignKeys) != 2 {
		t.Fatalf("ForeignKeys count = %d, want 2: %v", len(posts.ForeignKeys), posts.ForeignKeys)
	}
	haveInline, haveConstraint := false, false
	for _, fk := range posts.ForeignKeys {
		if fk.RefTable != "users" || fk.RefColumn != "id" {
			t.Errorf("unexpected fk: %+v", fk)
		}
		if fk.Column == "user_id" {
			haveInline = true
		}
		if fk.Column == "editor_id" {
			haveConstraint = true
		}
	}
	if !haveInline {
		t.Errorf("inline FK on user_id missing")
	}
	if !haveConstraint {
		t.Errorf("named CONSTRAINT FK missing")
	}
}
