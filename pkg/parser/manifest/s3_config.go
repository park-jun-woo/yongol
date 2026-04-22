//ff:type feature=projectconfig type=model
//ff:what S3 스토리지 설정 구조체
package manifest

type S3Config struct {
	Bucket string `yaml:"bucket"`
	Region string `yaml:"region"`
}
