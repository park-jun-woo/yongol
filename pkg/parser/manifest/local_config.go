//ff:type feature=projectconfig type=model
//ff:what 로컬 파일 스토리지 설정 구조체
package manifest

type LocalConfig struct {
	Root string `yaml:"root"`
}
