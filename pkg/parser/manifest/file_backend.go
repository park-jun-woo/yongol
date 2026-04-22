//ff:type feature=projectconfig type=model
//ff:what 파일 스토리지 백엔드 설정 구조체
package manifest

// FileBackend configures file storage backend (s3 | local).
type FileBackend struct {
	Backend string       `yaml:"backend"` // "s3" or "local"
	S3      *S3Config    `yaml:"s3"`
	Local   *LocalConfig `yaml:"local"`
}
