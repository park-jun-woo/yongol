//ff:func feature=validate type=test-helper control=sequence topic=manifest-infra
//ff:what usersTableFixture — XDN 테스트용 users 테이블 컬럼 셋 (id/email/role/user_id + 타입 변종)

package manifest_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// usersTableFixture builds a deterministic ddl.Table used by the XDN-03
// and XDN-04 test suites. It mixes correctly-typed columns (id, email,
// role, user_id, org_id, is_admin) with one intentional type mismatch
// (profile_id is bool) so both happy-path and negative tests share the
// same fixture.
func usersTableFixture() ddl.Table {
	return ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":         {Name: "id", RawType: "BIGINT"},
			"user_id":    {Name: "user_id", RawType: "BIGINT"},
			"org_id":     {Name: "org_id", RawType: "BIGINT"},
			"email":      {Name: "email", RawType: "TEXT"},
			"role":       {Name: "role", RawType: "TEXT"},
			"is_admin":   {Name: "is_admin", RawType: "BOOLEAN"},
			"profile_id": {Name: "profile_id", RawType: "BOOLEAN"},
		},
	}
}
