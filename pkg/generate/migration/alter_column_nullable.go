//ff:type feature=migration type=model
//ff:what AlterColumnNullable — ALTER COLUMN SET/DROP NOT NULL Operation
package migration

type AlterColumnNullable struct {
	Table, Column string
	From, To      bool
	Backfill      string // from @backfill hint
}
