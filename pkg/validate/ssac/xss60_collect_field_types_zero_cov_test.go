//ff:func feature=validate type=test control=sequence
//ff:what TestXss60ByName_ZeroCov — XSS-60 publish/subscribe 타입 수집·비교 헬퍼 직접 호출
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60CollectFieldTypes_ZeroCov(t *testing.T) {
	fields := map[string]string{}
	seq := parsessac.Sequence{
		Type:  "publish",
		Topic: "order.completed",
		Inputs: map[string]string{
			"status": `"done"`, // string literal
			"count":  "42",     // int literal
		},
		Fields: map[string]string{
			"note": `"x"`,
		},
	}
	fn := parsessac.ServiceFunc{Name: "Pub"}
	xss60CollectFieldTypes(fields, seq, fn, map[string]*ddl.Table{})
	if fields["status"] != "string" || fields["count"] != "int64" || fields["note"] != "string" {
		t.Errorf("collected field types = %v", fields)
	}
}
