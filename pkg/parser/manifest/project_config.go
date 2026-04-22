//ff:type feature=projectconfig type=model
//ff:what manifest.yaml 프로젝트 설정 최상위 구조체
package manifest

// ProjectConfig represents the manifest.yaml project configuration.
type ProjectConfig struct {
	APIVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   Metadata        `yaml:"metadata"`
	Backend    Backend         `yaml:"backend"`
	Frontend   Frontend        `yaml:"frontend"`
	Deploy     Deploy          `yaml:"deploy"`
	Session    *BuiltinBackend `yaml:"session"`
	Cache      *BuiltinBackend `yaml:"cache"`
	File       *FileBackend    `yaml:"file"`
	Queue      *QueueBackend   `yaml:"queue"`
	Authz      *AuthzConfig    `yaml:"authz"`
}
