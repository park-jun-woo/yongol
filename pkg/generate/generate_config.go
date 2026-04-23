//ff:type feature=generate type=model
//ff:what generateConfig — Generate 옵션 누적 구조체
package generate

type generateConfig struct {
	migration MigrationHook
}
