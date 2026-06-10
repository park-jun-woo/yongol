//ff:type feature=stml-parse type=model
//ff:what data-action 요소와 하위 필드를 나타내는 구조체
package stml

// ActionBlock represents a data-action element and its descendant fields.
type ActionBlock struct {
	Tag         string        // original HTML tag
	ClassName   string        // class attribute value
	OperationID string        // data-action value (e.g. "CreateReservation")
	Params      []ParamBind   // data-param-* attributes on this element
	Fields      []FieldBind   // descendant data-field attributes (for validation)
	Children    []ChildNode   // all children in DOM order (for codegen)
	SubmitText  string        // text of button[type=submit]
	EnabledWhen string        // data-enabled-when guard condition (empty if unset)
	Invalidates []string      // data-invalidates operationIds to refetch on success (empty if unset)
	CaptureRaw  string        // raw data-capture attribute value (TM-20 re-parses it for syntax diagnostics)
	Captures    []CaptureBind // parsed data-capture bindings (empty when absent or syntactically invalid)
	Redirect    string        // data-redirect static path navigated to on action success (empty if unset)
	OnErrorNode bool          // true when a descendant element carries data-on-error

	// RowMutateArg is codegen-populated (pkg/generate/react/stml, like
	// EachBlock.KeyField): the call-site mutate() argument object for a row
	// action inside data-each whose params use item.<Field> sources. The
	// parser always leaves it empty.
	RowMutateArg string
}
