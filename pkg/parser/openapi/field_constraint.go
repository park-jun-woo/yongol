//ff:type feature=manifest type=model
//ff:what FieldConstraint — OpenAPI 스키마 property의 제약조건
package openapi

// FieldConstraint holds constraints for a single OpenAPI schema property.
type FieldConstraint struct {
	Type      string
	Format    string
	MaxLength *int
	MinLength *int
	Enum      []string
	Required  bool
	// Line: 해당 property 가 선언된 줄 번호 (1-based, 0 = 미상).
	// LineIndex 와 매칭되어 채워지며, 검증 진단의 file:line 정확성 확보용.
	Line int
}
