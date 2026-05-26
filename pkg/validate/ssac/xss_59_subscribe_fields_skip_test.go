//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-59 skip — non-subscribe/nil-param/missing-struct/empty-topic/fields-key/no-publish skip

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss59SubscribeFieldsSkip(t *testing.T) {
	t.Run("skips_non_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Get", FileName: "s/g.ssac"}}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_nil_param", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "Sub", FileName: "s/sub.ssac", Subscribe: &parsessac.SubscribeInfo{Topic: "event"},
		}}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_missing_struct", func(t *testing.T) {
		pub := parsessac.ServiceFunc{Name: "Pub", FileName: "s/p.ssac", Sequences: []parsessac.Sequence{{Type: "publish", Topic: "event", Inputs: map[string]string{"ID": "x.ID"}, Line: 2}}}
		sub := parsessac.ServiceFunc{Name: "Sub", FileName: "s/s.ssac", Subscribe: &parsessac.SubscribeInfo{Topic: "event"}, Param: &parsessac.ParamInfo{TypeName: "Missing", VarName: "message"}}
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{pub, sub}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("publish_with_fields_key", func(t *testing.T) {
		pub := parsessac.ServiceFunc{Name: "Pub", FileName: "s/p.ssac", Sequences: []parsessac.Sequence{{Type: "publish", Topic: "event", Fields: map[string]string{"OrderID": "order.ID"}, Line: 2}}}
		sub := parsessac.ServiceFunc{Name: "Sub", FileName: "s/s.ssac", Subscribe: &parsessac.SubscribeInfo{Topic: "event"}, Param: &parsessac.ParamInfo{TypeName: "Msg", VarName: "message"}, Structs: []parsessac.StructInfo{{Name: "Msg", Fields: []parsessac.StructField{{Name: "OrderID", Type: "int64"}}}}}
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{pub, sub}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_empty_topic_in_scan", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Bad", FileName: "s/bad.ssac", Sequences: []parsessac.Sequence{{Type: "publish", Topic: "", Inputs: map[string]string{"ID": "x.ID"}, Line: 2}}}}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_no_publish_keys", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Sub", FileName: "s/s.ssac", Subscribe: &parsessac.SubscribeInfo{Topic: "event"}, Param: &parsessac.ParamInfo{TypeName: "Msg", VarName: "message"}, Structs: []parsessac.StructInfo{{Name: "Msg", Fields: []parsessac.StructField{{Name: "ID", Type: "int64"}}}}}}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
