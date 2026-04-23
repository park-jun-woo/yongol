//ff:type feature=migration type=model
//ff:what Column — 정규화된 단일 컬럼 정의
package migration

// Column is a single column definition, already normalised.
type Column struct {
	Name     string
	Type     CanonicalType
	RawType  string // original DDL type literal (e.g. "BIGSERIAL", "BIGINT"); used by D-8
	Nullable bool
	Default  string        // canonical default expression, "" = none
	Identity *IdentitySpec // GENERATED [ALWAYS|BY DEFAULT] AS IDENTITY, nil when absent
	Comment  string        // trailing -- comment (@sensitive / @nullable are consumed separately)
}
