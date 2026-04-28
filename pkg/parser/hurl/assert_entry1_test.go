//ff:func feature=crosscheck type=test-helper control=sequence topic=scenario-check
//ff:what assertEntry1 — ParseFile BodyFieldsCapturesAsserts 2번째 entry 검증

package hurl

import "testing"

// assertEntry1 verifies basic method/path and the Authorization header
// on the second entry parsed by the Phase002 extension fixture.
func assertEntry1(t *testing.T, e1 HurlEntry) {
	t.Helper()
	if e1.Method != "GET" || e1.Path != "/workflows/{{id}}" {
		t.Fatalf("entry[1] basics = %+v", e1)
	}
	if !hasHeader(e1.Headers, "Authorization") {
		t.Fatalf("entry[1] headers = %+v", e1.Headers)
	}
}
