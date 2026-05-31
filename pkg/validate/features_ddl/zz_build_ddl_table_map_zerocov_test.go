//ff:func feature=validate type=test control=sequence
//ff:what TestBuildDDLTableMap_ZeroCov — DDL 테이블 슬라이스 → 이름 맵 변환 직접 호출

package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestBuildDDLTableMap_ZeroCov(t *testing.T) {
	tables := []ddl.Table{{Name: "users"}, {Name: "orders"}}
	m := buildDDLTableMap(tables)
	if len(m) != 2 {
		t.Fatalf("map size = %d", len(m))
	}
	if m["users"] == nil || m["users"].Name != "users" {
		t.Errorf("users entry = %v", m["users"])
	}
	if m["orders"] == nil || m["orders"].Name != "orders" {
		t.Errorf("orders entry = %v", m["orders"])
	}
}
