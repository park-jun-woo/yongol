//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateHurlFromFeatures — 다중 feature stub 생성 / 잘못된 path 에러 / write 에러 분기 검증
package cliinit

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateHurlFromFeatures_WriteError(t *testing.T) {
	// specs/tests dir doesn't exist -> write fails.
	target := t.TempDir()
	feats := []features.Feature{{Op: "X", Path: "GET /x"}}
	if err := generateHurlFromFeatures(target, feats); err == nil {
		t.Fatal("want write error when dest dir missing")
	}
}
