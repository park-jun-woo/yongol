//ff:type feature=external type=model
//ff:what 메서드 정보 타입 정의
package external

type methodInfo struct {
	Name       string
	HTTPMethod string
	Path       string
	Params     []paramInfo
	ReturnType string // empty = error only
}
