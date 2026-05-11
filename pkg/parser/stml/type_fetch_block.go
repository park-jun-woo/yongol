//ff:type feature=stml-parse type=model
//ff:what data-fetch 요소와 하위 바인딩을 나타내는 구조체
package stml

// FetchBlock represents a data-fetch element and its descendant bindings.
type FetchBlock struct {
	Tag           string         // original HTML tag (e.g. "section")
	ClassName     string         // class attribute value
	OperationID   string         // data-fetch value (e.g. "ListMyReservations")
	Params        []ParamBind    // data-param-* attributes on this element
	Binds         []FieldBind    // descendant data-bind attributes (for validation)
	Eaches        []EachBlock    // descendant data-each blocks (for validation)
	States        []StateBind    // descendant data-state attributes (for validation)
	Components    []ComponentRef // descendant data-component attributes (for validation)
	Children      []ChildNode    // all children in DOM order (for codegen)
	NestedFetches []FetchBlock   // nested data-fetch blocks (for validation)
}
