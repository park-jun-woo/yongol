//ff:func feature=stml-gen type=generator control=sequence
//ff:what 기본 React/TSX Target을 반환한다
package stml

// DefaultTarget returns the built-in React/TSX target.
func DefaultTarget() Target {
	return &ReactTarget{}
}
