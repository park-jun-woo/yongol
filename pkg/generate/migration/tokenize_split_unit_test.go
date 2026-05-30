//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestTokenizeColumnDef(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{"simple", "id INTEGER NOT NULL", []string{"id", "INTEGER", "NOT", "NULL"}},
		{"paren preserved", "amount NUMERIC(10, 2) NOT NULL", []string{"amount", "NUMERIC(10, 2)", "NOT", "NULL"}},
		{"single quoted default", "status TEXT DEFAULT 'active'", []string{"status", "TEXT", "DEFAULT", "'active'"}},
		{"escaped single quote", "msg TEXT DEFAULT 'it''s'", []string{"msg", "TEXT", "DEFAULT", "'it''s'"}},
		{"double quoted ident", `"Order" INTEGER`, []string{`"Order"`, "INTEGER"}},
		{"nested parens", "v INTEGER CHECK (v IN (1, 2, 3))", []string{"v", "INTEGER", "CHECK", "(v IN (1, 2, 3))"}},
		{"empty", "", nil},
		{"leading whitespace collapses", "  a  b  ", []string{"a", "b"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := tokenizeColumnDef(c.in); !reflect.DeepEqual(got, c.want) {
				t.Errorf("tokenizeColumnDef(%q) = %#v, want %#v", c.in, got, c.want)
			}
		})
	}
}

func TestSplitStatements(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int // number of non-empty trimmed statements expected via spot-check below
	}{
		{"two statements", "CREATE TABLE a (id int); CREATE TABLE b (id int);", 2},
		{"semicolon in string is ignored", "INSERT INTO t VALUES ('a;b');", 1},
		{"semicolon in parens at top", "CREATE TABLE a (id int, c CHECK (x > 0));", 1},
		{"line comment stripped", "-- comment with ; inside\nSELECT 1;", 1},
		{"block comment with semicolon", "/* a; b */ SELECT 1;", 1},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := splitStatements(c.in)
			n := 0
			for _, s := range got {
				if len(trimSpaceSimple(s)) > 0 {
					n++
				}
			}
			if n != c.want {
				t.Errorf("splitStatements(%q) -> %d non-empty stmts, want %d: %#v", c.in, n, c.want, got)
			}
		})
	}
}

func trimSpaceSimple(s string) string {
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\n' || s[start] == '\t' || s[start] == '\r') {
		start++
	}
	end := len(s)
	for end > start && (s[end-1] == ' ' || s[end-1] == '\n' || s[end-1] == '\t' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}

func TestSplitTopLevel(t *testing.T) {
	cases := []struct {
		name string
		in   string
		sep  byte
		want []string
	}{
		{"simple comma", "a, b, c", ',', []string{"a", " b", " c"}},
		{"comma inside parens ignored", "a, b(1, 2), c", ',', []string{"a", " b(1, 2)", " c"}},
		{"comma inside single quote ignored", "a, 'x,y', c", ',', []string{"a", " 'x,y'", " c"}},
		{"comma inside double quote ignored", `a, "x,y", c`, ',', []string{"a", ` "x,y"`, " c"}},
		{"no sep", "abc", ',', []string{"abc"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := splitTopLevel(c.in, c.sep); !reflect.DeepEqual(got, c.want) {
				t.Errorf("splitTopLevel(%q,%q) = %#v, want %#v", c.in, c.sep, got, c.want)
			}
		})
	}
}

func TestParseColumnList(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{"parenthesised", "(a, b, c)", []string{"a", "b", "c"}},
		{"bare", "a, b", []string{"a", "b"}},
		{"uppercase lowered", "(ID, Name)", []string{"id", "name"}},
		{"quoted preserved", `("Order", id)`, []string{"Order", "id"}},
		{"empty entries skipped", "(a, , b)", []string{"a", "b"}},
		{"single", "id", []string{"id"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := parseColumnList(c.in); !reflect.DeepEqual(got, c.want) {
				t.Errorf("parseColumnList(%q) = %#v, want %#v", c.in, got, c.want)
			}
		})
	}
}

func TestCollectDefaultExpr(t *testing.T) {
	cases := []struct {
		name     string
		toks     []string
		wantExpr string
		wantN    int
	}{
		{"stops at NOT", []string{"0", "NOT", "NULL"}, "0", 1},
		{"function default", []string{"now()", "NOT", "NULL"}, "now()", 1},
		{"multi token expr stops at CHECK", []string{"'a'", "::", "text", "CHECK"}, "'a' :: text", 3},
		{"consumes all when no stop", []string{"42"}, "42", 1},
		{"empty", nil, "", 0},
		{"stops at GENERATED", []string{"x", "GENERATED"}, "x", 1},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			gotExpr, gotN := collectDefaultExpr(c.toks)
			if gotExpr != c.wantExpr || gotN != c.wantN {
				t.Errorf("collectDefaultExpr(%#v) = (%q,%d), want (%q,%d)", c.toks, gotExpr, gotN, c.wantExpr, c.wantN)
			}
		})
	}
}
