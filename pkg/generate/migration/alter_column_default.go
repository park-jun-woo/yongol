//ff:type feature=migration type=model
//ff:what AlterColumnDefault — ALTER COLUMN SET/DROP DEFAULT Operation
package migration

type AlterColumnDefault struct {
	Table, Column string
	From, To      string
}
