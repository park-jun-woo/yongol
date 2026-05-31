//ff:func feature=external type=test control=sequence
//ff:what TestStructTime — structHasTimeField/typesNeedTime time.Time 탐지 검증
package external

import (
	"testing"
)

func TestTypesNeedTime(t *testing.T) {
	none := []structType{
		{Name: "A", Fields: []structField{{GoType: "string"}}},
		{Name: "B", Fields: []structField{{GoType: "int"}}},
	}
	if typesNeedTime(none) {
		t.Error("expected false when no struct uses time.Time")
	}

	some := []structType{
		{Name: "A", Fields: []structField{{GoType: "string"}}},
		{Name: "B", Fields: []structField{{GoType: "time.Time"}}},
	}
	if !typesNeedTime(some) {
		t.Error("expected true when at least one struct uses time.Time")
	}

	if typesNeedTime(nil) {
		t.Error("expected false for nil slice")
	}
}
