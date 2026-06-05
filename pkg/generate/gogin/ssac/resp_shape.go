//ff:type feature=gen-gogin type=model topic=response
//ff:what respShape — 응답 래퍼 타입 1개의 분류 결과(종류 + 임베디드 타입명)

package ssac

// respShape records the classification of one response wrapper type.
type respShape struct {
	Kind respShapeKind
	// EmbeddedType is the name of the anonymously embedded type (== field
	// name) when Kind == shapeEmbedded. Empty for alias.
	EmbeddedType string
}
