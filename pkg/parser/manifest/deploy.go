//ff:type feature=projectconfig type=model
//ff:what 배포 설정 구조체
package manifest

type Deploy struct {
	Image  string `yaml:"image"`
	Domain string `yaml:"domain"`
}
