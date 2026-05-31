//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestFindOperationNilPaths(t *testing.T) {
	if got := findOperation(&openapi3.T{}, "GET", "/x"); got != nil {
		t.Errorf("expected nil for doc with nil Paths, got %v", got)
	}
}
