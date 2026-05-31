//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldSSaC — features 없음 0,nil / 기존파일 skip(count 0) / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSSaCNoFeatures(t *testing.T) {
	var out bytes.Buffer
	n, err := scaffoldSSaC(t.TempDir(), &features.FeaturesFile{}, "", nil, Config{}, &out)
	if err != nil || n != 0 {
		t.Fatalf("no features → %d, %v; want 0, nil", n, err)
	}
}
