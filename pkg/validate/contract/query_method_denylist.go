//ff:type feature=validate-contract type=model
//ff:what queryMethodDenylist — sqlc 런타임 메서드(WithTx 등) 제외 집합

package contract

// queryMethodDenylist lists methods present on the sqlc-generated
// `*Queries` receiver that are NOT driven by a named `-- name:` entry
// in any specs/db/*.sql file. PRV-02 must ignore them so the rule
// only ever flags real schema / query drift.
var queryMethodDenylist = map[string]bool{
	"WithTx": true,
}
