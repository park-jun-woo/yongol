//ff:func feature=manifest type=test control=sequence
//ff:what isFuncManagedAnnotation — `-- @func-managed` 주석 감지 (그 외는 false)

package ddl

import "testing"

func TestIsFuncManagedAnnotation(t *testing.T) {
	if !isFuncManagedAnnotation("-- @func-managed") {
		t.Errorf("missed -- @func-managed")
	}
	if isFuncManagedAnnotation("-- @archived") {
		t.Errorf("false positive for -- @archived")
	}
	if isFuncManagedAnnotation("-- some note") {
		t.Errorf("false positive")
	}
	if isFuncManagedAnnotation("CREATE TABLE x (") {
		t.Errorf("false positive for CREATE TABLE")
	}
}
