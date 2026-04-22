//ff:type feature=ssac-parse type=model
//ff:what 결과 바인딩 타입
package ssac

// Result는 결과 바인딩이다.
type Result struct {
	Type    string // "Course", "Reservation" (내부 타입)
	Var     string // "course", "reservations"
	Wrapper string // "Page", "Cursor", "" (제네릭 래퍼)
}
