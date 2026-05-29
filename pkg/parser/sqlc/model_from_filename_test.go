//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what modelFromFilename 테이블 테스트 — 복수형·snake_case → PascalCase 단수형
package sqlc

import "testing"

func TestModelFromFilename_Table(t *testing.T) {
	cases := []struct {
		filename string
		want     string
	}{
		{"users.sql", "User"},
		{"order_items.sql", "OrderItem"},
	}
	for _, tc := range cases {
		t.Run(tc.filename, func(t *testing.T) {
			got := modelFromFilename(tc.filename)
			if got != tc.want {
				t.Errorf("modelFromFilename(%q) = %q, want %q", tc.filename, got, tc.want)
			}
		})
	}
}
