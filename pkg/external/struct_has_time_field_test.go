//ff:func feature=external type=test control=sequence
//ff:what TestStructTime — structHasTimeField/typesNeedTime time.Time 탐지 검증
package external

import (
	"testing"
)

func TestStructHasTimeField(t *testing.T) {
	withTime := structType{Name: "A", Fields: []structField{
		{Name: "ID", GoType: "int64"},
		{Name: "CreatedAt", GoType: "time.Time"},
	}}
	if !structHasTimeField(withTime) {
		t.Error("expected true for struct containing time.Time field")
	}

	noTime := structType{Name: "B", Fields: []structField{
		{Name: "ID", GoType: "int64"},
		{Name: "Name", GoType: "string"},
	}}
	if structHasTimeField(noTime) {
		t.Error("expected false for struct without time.Time field")
	}

	empty := structType{Name: "C"}
	if structHasTimeField(empty) {
		t.Error("expected false for empty struct")
	}
}
