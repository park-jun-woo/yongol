//ff:func feature=gen-react type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 react.Generate 통합 커버리지
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func loadDummyFS_Integ(t *testing.T, root string) *yongol.Fullstack {
	t.Helper()
	det, err := yongol.DetectSSOTs(root)
	if err != nil {
		t.Fatalf("DetectSSOTs(%s): %v", root, err)
	}
	return yongol.ParseAll(root, det)
}
