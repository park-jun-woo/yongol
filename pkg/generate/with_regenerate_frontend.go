//ff:func feature=generate type=util control=sequence
//ff:what WithRegenerateFrontend — 프론트엔드 강제 재생성 플래그를 GenerateOption 으로 래핑
package generate

// WithRegenerateFrontend forces frontend regeneration even when the
// output directory already exists.
func WithRegenerateFrontend() GenerateOption {
	return func(c *generateConfig) { c.regenerateFrontend = true }
}
