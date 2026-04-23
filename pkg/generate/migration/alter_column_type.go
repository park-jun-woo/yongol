//ff:type feature=migration type=model
//ff:what AlterColumnType — ALTER COLUMN TYPE Operation
package migration

type AlterColumnType struct {
	Table, Column string
	From, To      CanonicalType
	Using         string // from @cast hint
}
