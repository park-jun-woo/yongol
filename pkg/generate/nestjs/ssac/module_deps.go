//ff:type feature=gen-nestjs type=model
//ff:what moduleDeps — NestJS module 렌더에 필요한 의존성 플래그 및 cross-feature 목록

package ssac

// moduleDeps holds the dependency flags and cross-feature module list needed to
// render a NestJS module file.
type moduleDeps struct {
	NeedsQueue           bool
	NeedsAuthz           bool
	NeedsSameFeatureStub bool
	CrossFeatures        []string // sorted cross-feature @call targets
}
