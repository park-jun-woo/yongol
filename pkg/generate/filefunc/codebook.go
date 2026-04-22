//ff:type feature=gen-filefunc type=model
//ff:what Codebook — filefunc codebook.yaml 최상위 구조 (required + optional)
package filefunc

// Codebook mirrors the filefunc codebook.yaml schema:
// required: feature + type (필수), optional: topic + ssot + pattern (선택).
type Codebook struct {
	Required Required `yaml:"required"`
	Optional Optional `yaml:"optional"`
}
