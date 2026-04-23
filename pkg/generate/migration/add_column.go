//ff:type feature=migration type=model
//ff:what AddColumn — ALTER TABLE ADD COLUMN Operation
package migration

type AddColumn struct {
	Table    string
	Column   *Column
	Backfill string // Phase004 — when set, prepend UPDATE before SET NOT NULL
}
