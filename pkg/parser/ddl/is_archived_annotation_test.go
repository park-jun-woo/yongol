//ff:func feature=manifest type=test control=sequence
//ff:what isArchivedAnnotation — `-- @archived` 주석 감지 (그 외는 false)

package ddl

import "testing"

func TestIsArchivedAnnotation(t *testing.T) {
	if !isArchivedAnnotation("-- @archived") {
		t.Errorf("missed -- @archived")
	}
	if isArchivedAnnotation("-- some note") {
		t.Errorf("false positive")
	}
	if isArchivedAnnotation("CREATE TABLE x (") {
		t.Errorf("false positive for CREATE TABLE")
	}
}
