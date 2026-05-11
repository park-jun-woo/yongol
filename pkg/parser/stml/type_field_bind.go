//ff:type feature=stml-parse type=model
//ff:what data-bind 또는 data-field 속성을 나타내는 구조체
package stml

// FieldBind represents a data-bind or data-field attribute.
type FieldBind struct {
	Name        string // field name (e.g. "Status", "Email")
	Tag         string // HTML tag (e.g. "span", "input")
	Type        string // input type attribute (e.g. "hidden", "number", "")
	ClassName   string // class attribute value
	Placeholder string // placeholder attribute value
}
