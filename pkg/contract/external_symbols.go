//ff:type feature=contract type=model
//ff:what ExternalSymbols — 함수 body 에서 참조한 외부 심볼(sqlc 쿼리·@call·struct field) 컨테이너

package contract

// ExternalSymbols lists the external references a preserved function
// makes. Precision is tuned for SSOT-driven drift detection: false
// positives are acceptable (they surface as informative noise during
// review) but a missed reference may hide a real contract break.
type ExternalSymbols struct {
	// SqlcQueries holds method names called on a `.Queries` receiver,
	// e.g. `server.Queries.GetUserByID(...)` contributes "GetUserByID".
	SqlcQueries []string
	// CallTargets holds `<pkg>.<Func>(...)` call expressions whose
	// package selector is a single identifier outside the callExprPkgDenylist
	// (so method calls on locals / receivers do not pollute the list).
	CallTargets []string
	// DDLFields holds `<recv>.<Field>` selector expressions where the
	// field name starts with an upper-case letter and the selector is
	// not used as a call target — these match the shape of exported
	// struct fields derived from DDL columns.
	DDLFields []string
}
