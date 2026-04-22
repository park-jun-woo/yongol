//ff:func feature=ssac-parse type=parser control=iteration dimension=1 topic=response
//ff:what @response 블록 내부 라인을 파싱하여 필드 맵 반환
package ssac

import "strings"

// parseResponseFields는 @response 블록 내부 라인을 파싱한다.
func parseResponseFields(lines []string) map[string]string {
	fields := make(map[string]string)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, ",")
		colonIdx := strings.IndexByte(line, ':')
		if colonIdx < 0 {
			continue
		}
		key := strings.TrimSpace(line[:colonIdx])
		val := strings.TrimSpace(line[colonIdx+1:])
		if key != "" && val != "" {
			fields[key] = val
		}
	}
	return fields
}
