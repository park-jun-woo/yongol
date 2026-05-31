//ff:func feature=validate type=test control=iteration dimension=1 topic=policy-check
//ff:what XDP-31 test — @ownership table must exist in DDL (Ground 기반)
package ddl_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithGroundTables(tables []string, policies []rego.Policy) *yongol.Fullstack {
	set := map[string]bool{}
	for _, t := range tables {
		set[t] = true
	}
	fs := &yongol.Fullstack{ParsedPolicies: policies}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"DDL.table": set}})
	return fs
}
