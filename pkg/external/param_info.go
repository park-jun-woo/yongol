//ff:type feature=external type=model
//ff:what 파라미터 정보 타입 정의
package external

type paramInfo struct {
	Name   string
	GoType string
	In     string // "body", "path"
}
