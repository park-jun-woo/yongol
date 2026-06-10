//ff:type feature=projectconfig type=model
//ff:what 프로젝트 메타데이터 구조체
package manifest

type Metadata struct {
	Name string `yaml:"name"`
	// Description is an optional human-readable note. The `yongol init`
	// scaffolder emits `metadata.description`, so it must be a recognised
	// field once strict decoding (KnownFields) is enabled in Load.
	Description string `yaml:"description,omitempty"`
}
