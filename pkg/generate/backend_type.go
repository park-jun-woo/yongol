//ff:type feature=generate type=model
//ff:what BackendType — 백엔드 코드 생성기 식별자
package generate

// BackendType identifies which backend generator to run.
type BackendType string

const (
	GoGin BackendType = "go-gin"
)
