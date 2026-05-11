//ff:type feature=stml-parse type=model
//ff:what 블록 내부의 자식 요소를 DOM 순서로 나타내는 구조체
package stml

// ChildNode represents any child element inside a block, preserving DOM order.
type ChildNode struct {
	Kind      string         // "bind", "each", "state", "component", "static", "action", "fetch"
	Bind      *FieldBind
	Each      *EachBlock
	State     *StateBind
	Component *ComponentRef
	Static    *StaticElement
	Action    *ActionBlock
	Fetch     *FetchBlock
}
