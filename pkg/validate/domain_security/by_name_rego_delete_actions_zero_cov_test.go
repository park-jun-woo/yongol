//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameRegoDeleteActions_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{{
			Rules: []rego.AllowRule{
				{Resource: "items", Actions: []string{"delete"}},
				{Resource: "users", Actions: []string{"read"}},
			},
		}},
	}
	result := collectRegoDeleteActions(fs)
	if _, ok := result["items"]; !ok {
		t.Errorf("collectRegoDeleteActions missing delete resource")
	}
}
