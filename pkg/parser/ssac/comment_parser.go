//ff:type feature=ssac-parse type=model
//ff:what 주석 파싱 상태 관리 타입
package ssac

// commentParser는 주석 파싱 상태를 관리한다.
type commentParser struct {
	sequences            []Sequence
	responseLines        []string
	inResponse           bool
	responseSuppressWarn bool
	responseStartLine    int // @response 첫 줄의 1-based 라인 번호
}
