//ff:type feature=features type=model
//ff:what TableDef — features.yaml tables 섹션의 단일 테이블 정의 (관계·상태)
package features

// TableDef represents a single table entry in the features.yaml tables section.
type TableDef struct {
	HasMany   []string `yaml:"has_many"`
	BelongsTo []string `yaml:"belongs_to"`
	States    []string `yaml:"states"`
}
