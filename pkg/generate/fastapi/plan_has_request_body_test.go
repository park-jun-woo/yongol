//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestPlanHasRequestBody — planHasRequestBody 메서드·바디 보유 분기 검증
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestPlanHasRequestBody(t *testing.T) {
	body := []ir.BodyFieldMeta{{Name: "title"}}
	cases := []struct {
		name   string
		method string
		fields []ir.BodyFieldMeta
		want   bool
	}{
		{"post with body", "post", body, true},
		{"put with body", "PUT", body, true},
		{"patch with body", "Patch", body, true},
		{"post no body", "POST", nil, false},
		{"get with body", "GET", body, false},
		{"delete with body", "DELETE", body, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			plan := &ir.ServicePlan{HTTPMethod: c.method, BodyFields: c.fields}
			if got := planHasRequestBody(plan); got != c.want {
				t.Errorf("planHasRequestBody(%q, %d fields) = %v, want %v",
					c.method, len(c.fields), got, c.want)
			}
		})
	}
}
