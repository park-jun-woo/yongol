//ff:func feature=generate type=util control=iteration dimension=1
//ff:what applyGenerateOptions — GenerateOption 슬라이스를 config 에 차례로 적용
package generate

func applyGenerateOptions(cfg *generateConfig, hooks []GenerateOption) {
	for _, h := range hooks {
		h(cfg)
	}
}
