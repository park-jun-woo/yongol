//ff:func feature=validate type=test control=iteration dimension=1 topic=authz-check
//ff:what xas60Ground -- XAS-60 테스트용 Ground 생성 헬퍼

package ssac_authz

import "github.com/park-jun-woo/yongol/pkg/rule"

func xas60Ground(fields ...string) *rule.Ground {
	set := rule.StringSet{}
	for _, f := range fields {
		set[f] = true
	}
	return &rule.Ground{
		Lookup: map[string]rule.StringSet{"Authz.checkRequest": set},
		Types:  map[string]string{},
		Pairs:  map[string]rule.StringSet{},
		Config: map[string]bool{},
		Vars:   rule.StringSet{},
		Flags:  rule.StringSet{},
	}
}
