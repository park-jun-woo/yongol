//ff:func feature=validate-contract type=test-helper control=sequence
//ff:what buildFSForPRV02 — PRV-02 테스트용 SQLc/DDL/Ground 세트 Fullstack 헬퍼

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFSForPRV02 assembles a Fullstack with a single sqlc query
// (UserFindByID), a users DDL table (id/email/created_at), and a
// Ground `SSaC.callRef` entry `billing.checkCredits` — enough surface
// to exercise all three PRV-02 symbol categories.
func buildFSForPRV02() *yongol.Fullstack {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "UserFindByID", Method: "FindByID"},
		},
		DDLTables: []ddl.Table{
			{Name: "users", Columns: map[string]ddl.Column{
				"id":         {Name: "id", RawType: "BIGINT"},
				"email":      {Name: "email", RawType: "VARCHAR(255)"},
				"created_at": {Name: "created_at", RawType: "TIMESTAMPTZ"},
			}},
		},
	}
	g := &rule.Ground{
		Lookup:  map[string]rule.StringSet{},
		Types:   map[string]string{},
		Pairs:   map[string]rule.StringSet{},
		Config:  map[string]bool{},
		Vars:    rule.StringSet{},
		Flags:   rule.StringSet{},
		Schemas: map[string][]string{},
	}
	g.Lookup["SSaC.callRef"] = rule.StringSet{"billing.checkCredits": true}
	fs.SetGround(g)
	return fs
}
