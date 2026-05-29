//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what collectSensitiveKeysFromTable — 한 테이블의 @sensitive 컬럼명을 seen 맵에 누적

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestCollectSensitiveKeysFromTable(t *testing.T) {
	cols := map[string]ddl.Column{
		"id":       {Name: "id"},
		"password": {Name: "password", Sensitive: true},
		"ssn":      {Name: "ssn", Sensitive: true},
	}
	seen := map[string]bool{"existing": true}
	collectSensitiveKeysFromTable(cols, seen)

	if !seen["password"] || !seen["ssn"] {
		t.Errorf("sensitive columns must be recorded, got %v", seen)
	}
	if seen["id"] {
		t.Errorf("non-sensitive column must not be recorded, got %v", seen)
	}
	if !seen["existing"] {
		t.Errorf("pre-existing keys must be preserved, got %v", seen)
	}
}
