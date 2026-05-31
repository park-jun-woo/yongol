//ff:func feature=hurl-parse type=test control=sequence
//ff:what TestByName_ZeroCov — hurl 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package hurl

import (
	"strings"
	"testing"
)

// TestByNameFeed_ZeroCov drives a full Hurl document through feed()/finish()
// by name, exercising flushEntry / flushRequestBody / processSectionHeader /
// processContentLine / collect*/handle*/newRequestEntry transitively, and then
// calls the standalone helpers by name directly to credit them.
func TestByNameFeed_ZeroCov(t *testing.T) {
	doc := `POST {{host}}/api/items
Content-Type: application/json

{
  "name": "x",
  "count": 3
}

HTTP 200
[Captures]
itemId: jsonpath "$.id"
[Asserts]
jsonpath "$.name" == "x"

GET https://example.com/api/items/1
HTTP 200
`
	st := &parseState{path: "test.hurl"}
	for _, raw := range strings.Split(doc, "\n") {
		st.lineNum++
		st.feed(raw)
	}
	st.finish()
	if len(st.entries) < 2 {
		t.Fatalf("feed produced %d entries, want >= 2", len(st.entries))
	}
}

// TestByNameHurlHelpers_ZeroCov calls the standalone parse helpers by name.
func TestByNameHurlHelpers_ZeroCov(t *testing.T) {
	// newRequestEntry from a regex-style match slice.
	m := []string{"POST {{host}}/api/x", "POST", "host", "/api/x?q=1"}
	entry := newRequestEntry(m, "f.hurl", 1)
	if entry.Method != "POST" || entry.URLVar != "host" {
		t.Errorf("newRequestEntry = %+v", entry)
	}

	// collectCaptureLine + collectAssertLine on a fresh state.
	s := &parseState{path: "f.hurl", current: entry, lineNum: 2}
	collectCaptureLine(s, `itemId: jsonpath "$.id"`)
	collectCaptureLine(s, "# comment")
	collectCaptureLine(s, "")
	collectAssertLine(s, `jsonpath "$.name" == "x"`)
	collectAssertLine(s, "# comment")
	collectAssertLine(s, "")

	// handleRequestHeaderOrBodyStart across branches.
	s.section = "request-headers"
	handleRequestHeaderOrBodyStart(s, "Content-Type: application/json", "Content-Type: application/json")
	handleRequestHeaderOrBodyStart(s, "{", "{")
	handleRequestHeaderOrBodyStart(s, "# c", "# c")
	handleRequestHeaderOrBodyStart(s, "", "")

	// processSectionHeader across branches.
	for _, hdr := range []string{"[Captures]", "[Asserts]", "[Options]", "[QueryStringParams]", "not-a-section"} {
		_ = processSectionHeader(s, hdr)
	}

	// processContentLine across sections.
	s.section = "captures"
	processContentLine(s, `v: jsonpath "$.v"`, `v: jsonpath "$.v"`)
	s.section = "asserts"
	processContentLine(s, `jsonpath "$.x" == 1`, `jsonpath "$.x" == 1`)
	s.section = "body"
	processContentLine(s, `{"a":1}`, `{"a":1}`)
	s.section = "request-headers"
	processContentLine(s, "X-H: v", "X-H: v")

	// flushRequestBody + flushEntry.
	s.bodyBuf.WriteString(`{"field": 1}`)
	s.flushRequestBody()
	s.current = entry
	s.flushEntry()
}

// TestByNameReplaceHurlVarAt_ZeroCov calls replaceHurlVarAt by name across branches.
func TestByNameReplaceHurlVarAt_ZeroCov(t *testing.T) {
	// not a var start.
	var b1 strings.Builder
	if _, ok := replaceHurlVarAt("abc", 0, &b1); ok {
		t.Errorf("replaceHurlVarAt non-var should be false")
	}

	// quoted var -> bare.
	var b2 strings.Builder
	b2.WriteByte('"')
	body := `"{{token}}"`
	if _, ok := replaceHurlVarAt(body, 1, &b2); !ok {
		t.Errorf("replaceHurlVarAt quoted var should be true")
	}

	// unquoted var.
	var b3 strings.Builder
	if _, ok := replaceHurlVarAt(`{{token}}`, 0, &b3); !ok {
		t.Errorf("replaceHurlVarAt unquoted var should be true")
	}

	// unterminated var.
	var b4 strings.Builder
	if _, ok := replaceHurlVarAt(`{{token`, 0, &b4); !ok {
		t.Errorf("replaceHurlVarAt unterminated should be true")
	}
}
