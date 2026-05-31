//ff:func feature=validate type=test control=sequence
//ff:what TestXss60ByName_ZeroCov — XSS-60 publish/subscribe 타입 수집·비교 헬퍼 직접 호출
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss60CollectPublishTypes_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
		Name: "Pub",
		Sequences: []parsessac.Sequence{
			{Type: "publish", Topic: "t.one", Inputs: map[string]string{"a": "1"}},
			{Type: "post", Topic: ""}, // skipped (not publish)
		},
	}}}
	out := xss60CollectPublishTypes(fs, map[string]*ddl.Table{})
	if out["t.one"]["a"] != "int64" {
		t.Errorf("publish types = %v", out)
	}
}
