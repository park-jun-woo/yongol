//ff:type feature=validate type=model topic=ssac-sqlc
//ff:what ReturningShape — sqlc RETURNING 절 형태 분류 (XQS-20 입력)

package ssac_sqlc

// ReturningShape describes the form of a sqlc query's RETURNING clause as
// seen by the cross-check XQS-20.
//
//   - ShapeNone    : no RETURNING clause (plain SELECT, :exec, etc.)
//   - ShapeFull    : `RETURNING *` or every DDL column listed → sqlc emits the model.
//   - ShapePartial : a strict subset of DDL columns → sqlc emits `<QueryName>Row`.
type ReturningShape string

// ReturningShape constants.
const (
	ShapeNone    ReturningShape = "none"
	ShapeFull    ReturningShape = "full"
	ShapePartial ReturningShape = "partial"
)
