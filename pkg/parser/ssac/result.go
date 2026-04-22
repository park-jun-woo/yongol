//ff:type feature=ssac-parse type=model
//ff:what Result — type representing a result binding
package ssac

// Result represents a result binding.
type Result struct {
	Type    string // "Course", "Reservation" (concrete type)
	Var     string // "course", "reservations"
	Wrapper string // "Page", "Cursor", "" (generic wrapper)
}
