//ff:func feature=parse-sqlc type=test control=sequence
//ff:what appendBodyLine test — 첫 줄 대입과 후속 줄 개행 연결 분기 검증

package sqlc

import "testing"

func TestAppendBodyLine(t *testing.T) {
	q := &QuerySpec{}
	appendBodyLine(q, "SELECT 1")
	if q.Body != "SELECT 1" {
		t.Fatalf("first line: got %q, want %q", q.Body, "SELECT 1")
	}
	appendBodyLine(q, "FROM t")
	if q.Body != "SELECT 1\nFROM t" {
		t.Fatalf("second line: got %q, want %q", q.Body, "SELECT 1\nFROM t")
	}
}
