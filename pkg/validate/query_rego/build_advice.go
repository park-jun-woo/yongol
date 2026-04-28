//ff:func feature=validate type=util control=sequence topic=policy-check
//ff:what buildAdvice — @ownership 누락 쿼리의 copy-paste 용 sqlc 쿼리 스텁

package query_rego

import "fmt"

// buildAdvice renders a ready-to-paste sqlc query stub. `via` annotations
// produce a JOIN form; direct mappings produce a single-table SELECT.
func buildAdvice(queryName, table, column, joinTable, joinFK string) string {
	if joinTable != "" && joinFK != "" {
		return fmt.Sprintf(
			"Add the following sqlc query to specs/db/queries/%s.sql:\n"+
				"-- name: %s :one\n"+
				"SELECT c.%s FROM %s c JOIN %s l ON l.%s = c.id WHERE l.id = @id;",
			joinTable, queryName, column, table, joinTable, joinFK)
	}
	return fmt.Sprintf(
		"Add the following sqlc query to specs/db/queries/%s.sql:\n"+
			"-- name: %s :one\n"+
			"SELECT %s FROM %s WHERE id = @id;",
		table, queryName, column, table)
}
