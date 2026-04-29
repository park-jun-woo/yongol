//ff:func feature=validate type=util control=selection topic=ssac-sqlc
//ff:what expectedSsacReturnType — RETURNING shape + Model + QueryName → 기대 SSaC 선언 타입

package ssac_sqlc

// expectedSsacReturnType reports the SSaC return-type identifier the user
// must declare for a sqlc query, given the query's RETURNING shape.
//
//   - ShapeFull    → the model name (sqlc emits the model directly)
//   - ShapePartial → "<QueryName>Row" (sqlc auto-generates a row struct)
//   - ShapeNone    → the model name (a SELECT with no RETURNING projects rows
//                    of the model when columns match; XQS-20 still permits a
//                    Row declaration if the SELECT lists a strict subset, but
//                    the *expected* canonical type for the no-RETURNING path
//                    remains the model)
//
// Returns "" when modelName is empty (defensive — callers should already have
// filtered).
func expectedSsacReturnType(shape ReturningShape, modelName, queryName string) string {
	if modelName == "" {
		return ""
	}
	switch shape {
	case ShapePartial:
		return queryName + "Row"
	default:
		return modelName
	}
}
