//ff:func feature=gen-react type=model control=sequence
//ff:what uiPrimitives — filename → shadcn-like primitive source map (Button/Card/Input 등 10개)

package react

// uiPrimitives returns the filename → source map for every emitted primitive.
// Order is insignificant (map) but file count (10) is asserted by Phase003
// verification. Keep keys sorted alphabetically in source for grep-friendly
// diffs; Go's map randomization affects iteration but not emission.
func uiPrimitives() map[string]string {
	return map[string]string{
		"Button.tsx":   primitiveButton,
		"Card.tsx":     primitiveCard,
		"Input.tsx":    primitiveInput,
		"Form.tsx":     primitiveForm,
		"Modal.tsx":    primitiveModal,
		"Table.tsx":    primitiveTable,
		"Select.tsx":   primitiveSelect,
		"Checkbox.tsx": primitiveCheckbox,
		"Badge.tsx":    primitiveBadge,
		"Tooltip.tsx":  primitiveTooltip,
	}
}
