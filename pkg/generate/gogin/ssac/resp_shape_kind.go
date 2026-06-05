//ff:type feature=gen-gogin type=model topic=response
//ff:what respShapeKind — oapi-codegen 응답 래퍼 형태(alias/embedded) 열거형

package ssac

// respShapeKind enumerates the two oapi-codegen response wrapper shapes.
type respShapeKind int

const (
	// shapeAlias is `type X SchemaType` — the wrapper aliases a schema and
	// exposes that schema's fields directly (Error/Code).
	shapeAlias respShapeKind = iota
	// shapeEmbedded is `type X struct{ EmbeddedType }` — the wrapper embeds
	// another response type; the embedded type's name doubles as the field
	// name (anonymous embed).
	shapeEmbedded
)
