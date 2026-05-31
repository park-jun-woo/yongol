//ff:func feature=validate type=test control=sequence topic=sqlc
//ff:what TestXQS15InputKeyInitialism — @call 입력 키의 Go initialism 위반 검출
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS15InputKeyInitialism_NoFuncs_ZeroCov(t *testing.T) {
	if diags := xqs15InputKeyInitialism(&yongol.Fullstack{}); len(diags) != 0 {
		t.Errorf("expected 0 diags for empty fullstack, got %d", len(diags))
	}
}
