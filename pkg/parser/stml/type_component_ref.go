//ff:type feature=stml-parse type=model
//ff:what data-component 속성을 나타내는 구조체
package stml

// ComponentRef represents a data-component attribute.
type ComponentRef struct {
	Name      string // component name (e.g. "DatePicker")
	Bind      string // data-bind value if present
	Field     string // data-field value if present
	ClassName string // class attribute value
}
