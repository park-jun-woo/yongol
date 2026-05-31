//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateRegoFromFeatures — allow rule stub 생성 / write 에러 분기 검증
package cliinit

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateRegoFromFeatures_WriteError(t *testing.T) {
	target := t.TempDir() // no specs/policy -> write fails
	feats := []features.Feature{{Op: "X", Path: "GET /x"}}
	if err := generateRegoFromFeatures(target, feats); err == nil {
		t.Fatal("want write error when dest dir missing")
	}
}
