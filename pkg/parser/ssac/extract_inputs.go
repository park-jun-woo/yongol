//ff:func feature=ssac-parse type=util control=sequence
//ff:what 문자열에서 {…} 블록을 추출하고 나머지 반환
package ssac

import "strings"

// extractInputs는 문자열에서 {…} 블록을 추출하고 나머지를 반환한다.
func extractInputs(s string) (map[string]string, string, error) {
	openIdx := strings.IndexByte(s, '{')
	if openIdx < 0 {
		return map[string]string{}, s, nil
	}
	closeIdx := strings.IndexByte(s, '}')
	if closeIdx < 0 {
		return map[string]string{}, s, nil
	}
	inputStr := s[openIdx : closeIdx+1]
	rest := strings.TrimSpace(s[closeIdx+1:])
	inputs, err := parseInputs(inputStr)
	return inputs, rest, err
}
