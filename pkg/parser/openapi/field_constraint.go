//ff:type feature=manifest type=model
//ff:what FieldConstraint — constraints for a single OpenAPI schema property
package openapi

// FieldConstraint holds constraints for a single OpenAPI schema property.
type FieldConstraint struct {
	Type      string
	Format    string
	MaxLength *int
	MinLength *int
	Minimum   *float64
	Maximum   *float64
	Pattern   string
	Enum      []string
	Required  bool
	// ItemType: array 타입일 때 items 의 타입 (e.g., "string", "integer").
	// 비-array 이면 빈 문자열.
	ItemType string
	// Line: 해당 property 가 선언된 줄 번호 (1-based, 0 = 미상).
	// LineIndex 와 매칭되어 채워지며, 검증 진단의 file:line 정확성 확보용.
	Line int
}
