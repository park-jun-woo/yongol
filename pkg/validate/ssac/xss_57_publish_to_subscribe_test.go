//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-57 — orphan @publish fires ERROR, matched @publish passes, non-publish skips

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss57PublishToSubscribe(t *testing.T) {
	t.Run("fires_on_orphan_publish", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "CompleteOrder",
				FileName: "service/complete_order.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type:  "publish",
						Topic: "order.completed",
						Line:  5,
					},
				},
			}},
		}
		diags := xss57PublishToSubscribe(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag for orphan publish, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XSS-57]") {
			t.Errorf("expected XSS-57 prefix, got %q", diags[0].Message)
		}
	})

	t.Run("passes_on_matched_publish", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{
				{
					Name:     "CompleteOrder",
					FileName: "service/complete_order.ssac",
					Sequences: []parsessac.Sequence{
						{
							Type:  "publish",
							Topic: "order.completed",
							Line:  5,
						},
					},
				},
				{
					Name:      "OnOrderCompleted",
					FileName:  "service/on_order_completed.ssac",
					Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
				},
			},
		}
		diags := xss57PublishToSubscribe(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for matched publish, got %d", len(diags))
		}
	})

	t.Run("skips_non_publish", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "GetCourse",
				FileName: "service/get_course.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type: "get",
						Line: 3,
					},
				},
			}},
		}
		diags := xss57PublishToSubscribe(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for non-publish, got %d", len(diags))
		}
	})

	t.Run("skips_empty_topic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "WeirdPublish",
				FileName: "service/weird.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type:  "publish",
						Topic: "",
						Line:  3,
					},
				},
			}},
		}
		diags := xss57PublishToSubscribe(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for empty topic, got %d", len(diags))
		}
	})
}
