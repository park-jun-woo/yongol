//ff:type feature=features type=model
//ff:what Feature — features.yaml 내 단일 기능 항목 (op, path, desc)
package features

// Feature represents a single entry in features.yaml.
type Feature struct {
	Op   string `yaml:"op"`
	Path string `yaml:"path"`
	Desc string `yaml:"desc"`
	Line int    `yaml:"-"` // 1-based line number in the source file
}
