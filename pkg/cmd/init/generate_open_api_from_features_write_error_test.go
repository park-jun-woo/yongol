//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateOpenAPIFromFeatures — path 그룹핑 stub 생성 / 잘못된 path 에러 / write 에러 분기 검증
package cliinit

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateOpenAPIFromFeatures_WriteError(t *testing.T) {
	target := t.TempDir() // no specs/api -> write fails
	feats := []features.Feature{{Op: "X", Path: "GET /x"}}
	if err := generateOpenAPIFromFeatures(target, templateData{}, feats); err == nil {
		t.Fatal("want write error when dest dir missing")
	}
}
