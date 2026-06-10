//ff:type feature=stml-parse type=model
//ff:what data-each 요소와 반복 항목 바인딩을 나타내는 구조체
package stml

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// EachBlock represents a data-each element and its descendant bindings.
type EachBlock struct {
	Tag           string                  // original HTML tag (e.g. "ul")
	ClassName     string                  // class attribute value
	ItemTag       string                  // repeated item tag (e.g. "li")
	ItemClassName string                  // repeated item class
	Field         string                  // array field name (e.g. "reservations")
	Actions       []ActionBlock           // row-level data-action inside the loop (for validation; item.* param sources)
	Binds         []FieldBind             // data-bind inside the loop (for validation)
	States        []StateBind             // data-state inside the loop (for validation)
	Components    []ComponentRef          // data-component inside the loop (for validation)
	Children      []ChildNode             // item children in DOM order (for codegen)
	KeyField      string                  // item key field (e.g. "id"); empty → use index
	Diags         []diagnostic.Diagnostic // parse-time diagnostics collected inside data-each
}
