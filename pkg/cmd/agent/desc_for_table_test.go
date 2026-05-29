//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestDescForTable — 테이블과 일치하는 feature 의 desc 반환, 미일치 시 빈 문자열 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestDescForTable(t *testing.T) {
	lookup := map[string]features.Feature{
		"f1": {Table: "users", Desc: "User accounts"},
		"f2": {Table: "orders", Desc: "Customer orders"},
	}
	if got := descForTable(lookup, "orders"); got != "Customer orders" {
		t.Errorf("descForTable(orders) = %q, want %q", got, "Customer orders")
	}
	if got := descForTable(lookup, "missing"); got != "" {
		t.Errorf("descForTable(missing) = %q, want empty", got)
	}
}
