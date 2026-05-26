//ff:type feature=gen-ir type=model
//ff:what ResponseField -- IR 응답 본문의 단일 필드 매핑

package ir

// ResponseField describes one field in a service response body. Name is the
// JSON field name and Source is the Go variable / expression that provides the
// value.
type ResponseField struct {
	// Name is the JSON response field name (e.g. "course", "instructor_name").
	Name string

	// Source is the variable or expression providing the value (e.g. "course",
	// "instructor.Name").
	Source string
}
