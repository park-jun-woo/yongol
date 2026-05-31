//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldByName_ZeroCov — scaffoldOpenAPI / scaffoldSSaCFeature 직접 호출 (LLM 미사용 분기)
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldOpenAPI_NoFeatures_ZeroCov(t *testing.T) {
	var out bytes.Buffer
	// No features → early return (0, nil); LLMCallFunc never invoked.
	n, err := scaffoldOpenAPI(t.TempDir(), &features.FeaturesFile{}, nil, Config{}, &out)
	if err != nil || n != 0 {
		t.Fatalf("no features → %d, %v; want 0, nil", n, err)
	}
}
