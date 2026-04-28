//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what TestQ12PgtypeUuidOverride — Q-12 4 케이스 (no-uuid skip / both ok / missing nullable / both absent)

package query

import "testing"

// TestQ12PgtypeUuidOverride covers the four Q-12 scenarios via a table:
//
//   1. DDL has no UUID column                                      → 0 diags (rule skipped)
//   2. DDL has UUID + sqlc.yaml has both NULL/NOT NULL overrides   → 0 diags
//   3. DDL has UUID + sqlc.yaml only has nullable=true             → 1 diag (missing nullable=false)
//   4. DDL has UUID + sqlc.yaml has no overrides at all            → 1 diag, both sides reported
//
// Per Phase001 spec: case 4 collapses the two missing entries into a
// single diagnostic ("Q-12 = one rule, one message"). The advice block
// contains both YAML stanzas so the user pastes once. Each row is run by
// runQ12PgtypeUuidOverrideCase so this func stays within the Q4 PURE
// line budget.
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
