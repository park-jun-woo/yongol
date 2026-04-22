//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what {key: value, ...} 형식 입력을 파싱하여 맵 반환
package ssac

import (
	"fmt"
	"strings"
)

// parseInputs는 {key: value, ...} 형식을 파싱한다.
func parseInputs(s string) (map[string]string, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	s = strings.TrimSpace(s)
	if s == "" {
		return map[string]string{}, nil
	}
	result := make(map[string]string)
	for _, pair := range strings.Split(s, ",") {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		colonIdx := strings.IndexByte(pair, ':')
		if colonIdx < 0 {
			return nil, fmt.Errorf("%q는 유효하지 않은 입력 형식입니다. \"{Key: value}\" 형식을 사용하세요", pair)
		}
		key := strings.TrimSpace(pair[:colonIdx])
		val := strings.TrimSpace(pair[colonIdx+1:])
		if key != "" && val != "" {
			result[key] = val
		}
	}
	return result, nil
}
