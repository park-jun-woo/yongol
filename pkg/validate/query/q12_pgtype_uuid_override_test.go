//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what q12PgtypeUuidOverride — UUID override 검증 (nil fs/pass/누락) 검증
package query

import (
	"testing"
)

func TestQ12PgtypeUuidOverride(t *testing.T) {
	cases := []q12UuidTestCase{
		{
			name:      "no-uuid skips",
			ddl:       q12DDLNoUUID,
			sqlc:      q12SqlcNoOverrides,
			wantDiags: 0,
		},
		{
			name:      "both overrides pass",
			ddl:       q12DDLWithUUID,
			sqlc:      q12SqlcBothOverrides,
			wantDiags: 0,
		},
		{
			name:           "missing nullable=false",
			ddl:            q12DDLWithUUID,
			sqlc:           q12SqlcOnlyNullable,
			wantDiags:      1,
			wantMsgSubstrs: []string{"[Q-12]", "nullable=false"},
		},
		{
			name:           "both entries absent collapse to one diag",
			ddl:            q12DDLWithUUID,
			sqlc:           q12SqlcNoOverrides,
			wantDiags:      1,
			wantMsgSubstrs: []string{"[Q-12]", "nullable=false and nullable=true"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { runQ12PgtypeUuidOverrideCase(t, tc) })
	}
}
