//ff:func feature=stml-parse type=test control=sequence
//ff:what TestSplitRolesAttr — data-roles 분해(공백 트림/빈 항목 제외/빈 값 nil) 검증

package stml

import (
	"reflect"
	"testing"
)

func TestSplitRolesAttr(t *testing.T) {
	if got := splitRolesAttr(""); got != nil {
		t.Errorf("empty: want nil, got %v", got)
	}
	if got := splitRolesAttr("admin"); !reflect.DeepEqual(got, []string{"admin"}) {
		t.Errorf("single: got %v", got)
	}
	if got := splitRolesAttr(" admin , manager "); !reflect.DeepEqual(got, []string{"admin", "manager"}) {
		t.Errorf("trim: got %v", got)
	}
	if got := splitRolesAttr("admin,,manager,"); !reflect.DeepEqual(got, []string{"admin", "manager"}) {
		t.Errorf("empty entries: got %v", got)
	}
}
