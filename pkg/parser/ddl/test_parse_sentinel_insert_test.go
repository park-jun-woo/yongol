//ff:func feature=manifest type=test control=sequence
//ff:what TestIsSentinelAnnotation — `-- @sentinel` 라인 판정 기본 케이스

package ddl

import (
	"testing"
)

// TestIsSentinelAnnotation exercises the three basic shapes: exact match,
// near-miss (without @), and a non-comment line.
func TestIsSentinelAnnotation(t *testing.T) {
	if !isSentinelAnnotation("-- @sentinel") {
		t.Errorf("missed -- @sentinel")
	}
	if isSentinelAnnotation("-- sentinel") {
		t.Errorf("false positive without @")
	}
	if isSentinelAnnotation("INSERT INTO t VALUES (0);") {
		t.Errorf("false positive for non-comment")
	}
}
