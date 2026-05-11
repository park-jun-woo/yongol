//ff:type feature=stml-parse type=model
//ff:what data-action 요소와 하위 필드를 나타내는 구조체
package stml

// ActionBlock represents a data-action element and its descendant fields.
type ActionBlock struct {
	Tag         string      // original HTML tag
	ClassName   string      // class attribute value
	OperationID string      // data-action value (e.g. "CreateReservation")
	Params      []ParamBind // data-param-* attributes on this element
	Fields      []FieldBind // descendant data-field attributes (for validation)
	Children    []ChildNode // all children in DOM order (for codegen)
	SubmitText  string      // text of button[type=submit]
}
