//ff:type feature=migration type=model
//ff:what InsertSentinel — @sentinel INSERT Operation (raw SQL body preserved)
package migration

// InsertSentinel wraps a verbatim `@sentinel` INSERT statement so it
// participates in the same phase-ordered Operation pipeline as the
// schema DDL operations. The Body field contains the full INSERT
// statement through the final `;`.
type InsertSentinel struct {
	Table string
	Body  string
}
