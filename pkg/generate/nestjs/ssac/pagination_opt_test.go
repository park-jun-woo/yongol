//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestPaginationOpt — pagination key/value → Prisma take/skip/cursor/passthrough 옵션 검증

package ssac

import "testing"

func TestPaginationOpt(t *testing.T) {
	tests := []struct {
		key  string
		val  string
		want string
	}{
		{"per_page", "v", "take: v"},
		{"limit", "v", "take: v"},
		{"page_offset", "v", "skip: v"},
		{"offset", "v", "skip: v"},
		{"cursor", "c", "cursor: c ? { id: c } : undefined"},
		{"custom", "x", "custom: x"},
	}
	for _, tt := range tests {
		if got := paginationOpt(tt.key, tt.val); got != tt.want {
			t.Errorf("paginationOpt(%q,%q) = %q, want %q", tt.key, tt.val, got, tt.want)
		}
	}
}
