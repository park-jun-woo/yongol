//ff:func feature=stml-parse type=util control=sequence dimension=1
//ff:what GuardRef를 dotted "model.field" 문자열로 반환한다
package stml

// Path returns the dotted "model.field" form of the reference.
func (r GuardRef) Path() string {
	return r.Model + "." + r.Field
}
