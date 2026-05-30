//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what parseHintLine — `@tag key=val` 파싱 / 비힌트 nil / 태그 소문자화

package ddl

import "testing"

func TestParseHintLine(t *testing.T) {
	t.Run("tag with args", func(t *testing.T) {
		h := parseHintLine("-- @Rename from=old to=new", "/x.sql", 3, "users", "email")
		if h == nil {
			t.Fatal("expected hint, got nil")
		}
		if h.Tag != "rename" {
			t.Errorf("Tag = %q, want rename", h.Tag)
		}
		if h.Args["from"] != "old" || h.Args["to"] != "new" {
			t.Errorf("Args = %v", h.Args)
		}
		if h.File != "/x.sql" || h.Line != 3 || h.TableCtx != "users" || h.ColumnCtx != "email" {
			t.Errorf("context = %+v", h)
		}
	})
	t.Run("not a hint comment", func(t *testing.T) {
		if h := parseHintLine("-- ordinary note", "", 1, "", ""); h != nil {
			t.Errorf("expected nil, got %+v", h)
		}
	})
	t.Run("at with no token", func(t *testing.T) {
		if h := parseHintLine("-- @", "", 1, "", ""); h != nil {
			t.Errorf("expected nil for bare @, got %+v", h)
		}
	})
	t.Run("arg without value ignored", func(t *testing.T) {
		h := parseHintLine("-- @cast bare type=int", "", 1, "", "")
		if h == nil {
			t.Fatal("nil")
		}
		if _, ok := h.Args["bare"]; ok {
			t.Errorf("bare token should not produce arg: %v", h.Args)
		}
		if h.Args["type"] != "int" {
			t.Errorf("type arg = %q", h.Args["type"])
		}
	})
}
