//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-58 — orphan @subscribe fires ERROR, matched @subscribe passes, non-subscribe skips

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss58SubscribeToPublish(t *testing.T) {
	t.Run("fires_on_orphan_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:      "OnOrderCompleted",
				FileName:  "service/on_order_completed.ssac",
				Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
				Line:      3,
			}},
		}
		diags := xss58SubscribeToPublish(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag for orphan subscribe, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XSS-58]") {
			t.Errorf("expected XSS-58 prefix, got %q", diags[0].Message)
		}
	})

	t.Run("passes_on_matched_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{
				{
					Name:     "CompleteOrder",
					FileName: "service/complete_order.ssac",
					Sequences: []parsessac.Sequence{
						{Type: "publish", Topic: "order.completed", Line: 5},
					},
				},
				{
					Name:      "OnOrderCompleted",
					FileName:  "service/on_order_completed.ssac",
					Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
				},
			},
		}
		diags := xss58SubscribeToPublish(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for matched subscribe, got %d", len(diags))
		}
	})

	t.Run("skips_non_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "GetCourse",
				FileName: "service/get_course.ssac",
			}},
		}
		diags := xss58SubscribeToPublish(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for non-subscribe, got %d", len(diags))
		}
	})

	t.Run("skips_empty_topic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:      "WeirdSub",
				FileName:  "service/weird_sub.ssac",
				Subscribe: &parsessac.SubscribeInfo{Topic: ""},
			}},
		}
		diags := xss58SubscribeToPublish(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for empty topic, got %d", len(diags))
		}
	})
}
