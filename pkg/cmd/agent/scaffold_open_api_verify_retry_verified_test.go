//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldOpenAPIVerifyRetry — verify 성공 / 미귀속 verify 에러 stop / 귀속 op 재시도+재조립 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldOpenAPIVerifyRetryVerified(t *testing.T) {
	doc := validOpenAPIDoc
	var offsets []pathOffset
	var out bytes.Buffer
	res := scaffoldOpenAPIVerifyRetry(&doc, &offsets, map[string]any{}, map[string][]string{},
		map[string]string{}, map[string]features.Feature{}, nil, false, 0, Config{}, &out,
		&features.FeaturesFile{})
	if !res.verified {
		t.Errorf("expected verified=true for a valid document, got %+v; out=%q", res, out.String())
	}
}
