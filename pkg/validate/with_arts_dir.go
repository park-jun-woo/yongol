//ff:func feature=validate type=util control=sequence
//ff:what WithArtsDir — contract drift 검증을 위한 생성 아티팩트 루트 디렉토리 주입

package validate

// WithArtsDir injects the artifacts directory (the second argument of
// `yongol generate`). When set to a non-empty path, the validate
// orchestrator also runs pkg/validate/contract.Run against that tree
// and merges its Diagnostics as a dedicated "contract" step result.
// When empty (the default), contract validation is skipped entirely
// — that is the correct behavior for `yongol validate <specs>` alone.
func WithArtsDir(dir string) Option {
	return func(c *config) { c.artsDir = dir }
}
