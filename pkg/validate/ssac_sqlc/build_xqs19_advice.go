//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what buildXqs19Advice — XQS-19 누락 쿼리 copy-paste advice

package ssac_sqlc

import "fmt"

// buildXqs19Advice composes the copy-paste advice. interface.yaml's
// canonical_queries is the reference source; the package name is used to
// direct the user to specs/db/queries/<pkg>.sql.
func buildXqs19Advice(pkg, query string) string {
	return fmt.Sprintf(
		"Add a sqlc query named %q to specs/db/queries/%s.sql. "+
			"Refer to ssac/pkg/%s/interface.yaml canonical_queries for the template.",
		query, pkg, pkg)
}
