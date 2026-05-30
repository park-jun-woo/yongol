//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what scanHintComments — 파일 라인별 hint 수집 (block-above / trailing / pending drain)

package ddl

import (
	"strings"
	"testing"
)

func TestScanHintComments(t *testing.T) {
	sql := `-- @allow_destructive
CREATE TABLE users (
  id BIGINT PRIMARY KEY,
  -- @rename from=old to=new
  new_col TEXT,
  email VARCHAR(255) -- @cast type=citext
);`
	out, err := scanHintComments(strings.NewReader(sql), "/users.sql")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	byTag := map[string]HintComment{}
	for _, h := range out {
		byTag[h.Tag] = h
	}

	// block-above hint attaches to the CREATE TABLE with BlockAbove.
	ad, ok := byTag["allow_destructive"]
	if !ok || ad.TableCtx != "users" || !ad.BlockAbove {
		t.Errorf("allow_destructive = %+v", ad)
	}
	// standalone hint drains onto the next DDL line (new_col).
	rn, ok := byTag["rename"]
	if !ok || rn.ColumnCtx != "new_col" || rn.TableCtx != "users" {
		t.Errorf("rename = %+v", rn)
	}
	if rn.Args["from"] != "old" || rn.Args["to"] != "new" {
		t.Errorf("rename args = %v", rn.Args)
	}
	// trailing hint attaches to its own line's column (email).
	ct, ok := byTag["cast"]
	if !ok || ct.ColumnCtx != "email" {
		t.Errorf("cast = %+v", ct)
	}
}

func TestScanHintComments_NoHints(t *testing.T) {
	sql := "CREATE TABLE t (id BIGINT);\n-- ordinary comment\n"
	out, err := scanHintComments(strings.NewReader(sql), "/t.sql")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(out) != 0 {
		t.Errorf("expected no hints, got %+v", out)
	}
}
