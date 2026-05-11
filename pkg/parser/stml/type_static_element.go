//ff:type feature=stml-parse type=model
//ff:what 바인딩이 없는 정적 HTML 요소를 나타내는 구조체
package stml

// StaticElement represents a non-binding HTML element to preserve in output.
type StaticElement struct {
	Tag       string      // HTML tag (e.g. "header", "h1")
	ClassName string      // class attribute value
	Text      string      // direct text content
	Children  []ChildNode // nested children
}
