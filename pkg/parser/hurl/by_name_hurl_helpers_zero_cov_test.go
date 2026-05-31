//ff:func feature=crosscheck type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — hurl 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package hurl

import (
	"testing"
)

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
