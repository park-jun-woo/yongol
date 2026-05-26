//ff:type feature=gen-ir type=model
//ff:what QueryMethod -- ServicePlan 이 사용하는 DB 쿼리 메서드 참조

package ir

// QueryMethod identifies a single database query method referenced by a
// ServicePlan.
type QueryMethod struct {
	// Name is the sqlc method name (e.g. "FindByID", "ListByCourseID").
	Name string

	// Package is the sqlc package prefix (e.g. "session"), empty if the
	// default queries package.
	Package string
}
