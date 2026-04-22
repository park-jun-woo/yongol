//ff:type feature=gen-filefunc type=model
//ff:what Required — codebook.yaml required 섹션 (feature + type 필수 키)
package filefunc

// Required section — 모든 //ff:func / //ff:type 어노테이션에 필수 키.
type Required struct {
	Feature map[string]string `yaml:"feature"`
	Type    map[string]string `yaml:"type"`
}
