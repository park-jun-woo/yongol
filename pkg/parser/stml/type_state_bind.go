//ff:type feature=stml-parse type=model
//ff:what data-state 속성을 나타내는 구조체
package stml

// StateBind represents a data-state attribute.
type StateBind struct {
	Tag       string      // original HTML tag (e.g. "p", "footer")
	ClassName string      // class attribute value
	Condition string      // condition expression (e.g. "reservations.empty", "canDelete")
	Text      string      // text content (e.g. "예약이 없습니다")
	Children  []ChildNode // nested children (e.g. action buttons inside state)
}
