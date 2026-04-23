//ff:type feature=migration type=model
//ff:what DropTable — DROP TABLE Operation
package migration

type DropTable struct {
	Name             string
	AllowDestructive bool // set by check_safety from hints
}
