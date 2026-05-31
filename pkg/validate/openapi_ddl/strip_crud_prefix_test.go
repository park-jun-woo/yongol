//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what stripCRUDPrefix — CRUD 동사 prefix 제거 + 비매칭 유지 검증

package openapi_ddl

import "testing"

func TestStripCRUDPrefix(t *testing.T) {
	tests := []struct {
		opID string
		want string
	}{
		{"getUser", "User"},
		{"listOrders", "Orders"},
		{"createWorkflow", "Workflow"},
		{"updateCategory", "Category"},
		{"deleteItem", "Item"},
		{"patchProfile", "Profile"},
		{"fetchData", "Data"},
		{"searchProducts", "Products"},
		{"findUser", "User"},
		{"doSomething", "doSomething"}, // no known prefix
		{"get", "get"},                 // too short (no rest)
		{"getuser", "getuser"},         // rest doesn't start with uppercase
		{"", ""},                       // empty
	}

	for _, tt := range tests {
		t.Run(tt.opID, func(t *testing.T) {
			got := stripCRUDPrefix(tt.opID)
			if got != tt.want {
				t.Errorf("stripCRUDPrefix(%q) = %q, want %q", tt.opID, got, tt.want)
			}
		})
	}
}
