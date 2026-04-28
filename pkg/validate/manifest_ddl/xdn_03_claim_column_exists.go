//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what XDN-03 — auth.claims 의 target column 이 user_table 에 존재하는지 검증

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdn03ClaimColumnExists checks every claim mapping
// (`<Field>: <col>[:<type>]`) against the user_table column set. One
// diagnostic per missing column; the claim line in manifest.yaml is
// surfaced when the parser captured it. Skipped when XDN-01 / XDN-02
// already would have fired (no auth, missing user_table field, or
// user_table not present in DDL).
func xdn03ClaimColumnExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !isAuthActive(fs) {
		return nil
	}
	auth := fs.Manifest.Backend.Auth
	if auth.UserTable == "" {
		return nil
	}
	tbl := findUserTable(fs, auth.UserTable)
	if tbl == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, field := range sortedClaimFields(auth.Claims) {
		def := auth.Claims[field]
		if _, ok := tbl.Columns[def.Key]; ok {
			continue
		}
		diags = append(diags, xdn03MissingColumnDiag(field, auth.UserTable, def))
	}
	return diags
}
