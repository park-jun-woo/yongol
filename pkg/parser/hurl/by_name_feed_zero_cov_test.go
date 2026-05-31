//ff:func feature=crosscheck type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — hurl 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package hurl

import (
	"strings"
	"testing"
)

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
