//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRouteSuffix — URLPath → FastAPI 라우트 suffix (prefix 제거)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRouteSuffix(t *testing.T) {
	cases := []struct {
		name string
		path string
		want string
	}{
		{"Empty", "", "/"},
		{"PrefixOnly", "/workflow", "/"},
		{"WithSuffix", "/workflow/:id", "/{id}"},
		{"DeepSuffix", "/workflow/:id/execute", "/{id}/execute"},
		{"NoLeadingSlashPrefixOnly", "workflow", "/"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			plan := &ir.ServicePlan{URLPath: c.path}
			if got := routeSuffix(plan); got != c.want {
				t.Errorf("routeSuffix(%q) = %q, want %q", c.path, got, c.want)
			}
		})
	}
}
