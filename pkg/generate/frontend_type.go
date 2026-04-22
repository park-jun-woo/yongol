//ff:type feature=generate type=model
//ff:what FrontendType — 프론트엔드 코드 생성기 식별자
package generate

// FrontendType identifies which frontend generator to run.
type FrontendType string

const (
	React FrontendType = "react"
)
