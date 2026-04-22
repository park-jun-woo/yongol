//ff:type feature=gen-filefunc type=model
//ff:what Optional — codebook.yaml optional 섹션 (topic / ssot / pattern 선택 키)
package filefunc

// Optional section — 해당할 때만 사용하는 키.
type Optional struct {
	Topic   map[string]string `yaml:"topic"`
	SSOT    map[string]string `yaml:"ssot,omitempty"`
	Pattern map[string]string `yaml:"pattern,omitempty"`
}
