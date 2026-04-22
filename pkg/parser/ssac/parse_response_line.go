//ff:func feature=ssac-parse type=parser control=sequence topic=response
//ff:what @response 줄을 파싱하여 Sequence 또는 멀티라인 시작 반환
package ssac

import "strings"

// parseResponseLine은 @response 줄을 파싱한다.
func parseResponseLine(line string) (*Sequence, bool, error) {
	tag := "@response"
	suppressWarn := false
	if strings.HasPrefix(line, "@response!") {
		tag = "@response!"
		suppressWarn = true
	}
	trimmed := strings.TrimSpace(strings.TrimPrefix(line, tag))
	if trimmed == "{" {
		return nil, true, nil
	}
	// 단일 행 구조체: @response { field: var, ... }
	if strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}") {
		inner := trimmed[1 : len(trimmed)-1]
		lines := strings.Split(inner, ",")
		return &Sequence{
			Type:         SeqResponse,
			Fields:       parseResponseFields(lines),
			SuppressWarn: suppressWarn,
		}, false, nil
	}
	// @response 간단쓰기: @response varName
	if trimmed != "" {
		return &Sequence{
			Type:         SeqResponse,
			Target:       trimmed,
			SuppressWarn: suppressWarn,
		}, false, nil
	}
	return nil, false, nil
}
