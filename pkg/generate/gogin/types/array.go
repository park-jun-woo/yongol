//ff:func feature=gen-gogin type=util control=sequence
//ff:what arrayBinding — TEXT[]/BIGINT[]/INTEGER[]/등 PG 배열 컬럼의 []T 매핑

package types

// arrayBinding returns the binding for a PG array column. The element
// kind is derived from elementHead (the head token after stripping the
// "[]" suffix in parseRawType). Only the four families with native sqlc
// support — string, integer, float, boolean — are produced as Go slices;
// anything else falls back to KindUnsupported via the dispatcher.
//
// PG arrays are themselves nullable (an empty array != NULL), so no
// pointer wrapping is added for nullable columns. Empty array literal is
// the conventional DEFAULT.
func arrayBinding(elementHead string, defaultLiteral string) GoTypeBinding {
	goElem, ok := arrayElementGoType(elementHead)
	return composeArrayBinding(goElem, ok, elementHead, defaultLiteral)
}
