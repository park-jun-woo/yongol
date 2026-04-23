//ff:func feature=external type=loader control=sequence
//ff:what readSource — URL 또는 파일 경로에서 OpenAPI 소스 데이터 읽기

package external

import (
	"os"
	"strings"
)

func readSource(source string) ([]byte, error) {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return readHTTPSource(source)
	}
	return os.ReadFile(source)
}
