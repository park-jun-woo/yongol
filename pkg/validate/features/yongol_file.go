//ff:type feature=validate type=model topic=features-structural
//ff:what yongolFile — specs/.yongol 파일 구조 (해시 맵)

package features

// yongolFile represents the structure of specs/.yongol.
type yongolFile struct {
	Hashes map[string]string `yaml:"hashes"`
}
