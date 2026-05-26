//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-38 — lowercase @call function fires ERROR, uppercase passes, non-call skips

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss38CallFuncLowercase(t *testing.T) {
	t.Run("fires_on_lowercase", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "BadCall",
				FileName: "service/bad_call.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type:  "call",
						Model: "auth.refreshToken",
						Line:  5,
					},
				},
			}},
		}
		diags := xss38CallFuncLowercase(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag for lowercase func, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XSS-38]") {
			t.Errorf("expected XSS-38 prefix, got %q", diags[0].Message)
		}
	})

	t.Run("passes_on_uppercase", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "GoodCall",
				FileName: "service/good_call.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type:  "call",
						Model: "auth.RefreshToken",
						Line:  5,
					},
				},
			}},
		}
		diags := xss38CallFuncLowercase(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for uppercase func, got %d", len(diags))
		}
	})

	t.Run("skips_non_call", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "GetCourse",
				FileName: "service/get_course.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type:  "get",
						Model: "Course.findByID",
						Line:  3,
					},
				},
			}},
		}
		diags := xss38CallFuncLowercase(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for non-call seq, got %d", len(diags))
		}
	})

	t.Run("skips_empty_method", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:     "WeirdCall",
				FileName: "service/weird_call.ssac",
				Sequences: []parsessac.Sequence{
					{
						Type:  "call",
						Model: "auth",
						Line:  3,
					},
				},
			}},
		}
		diags := xss38CallFuncLowercase(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags for empty method, got %d", len(diags))
		}
	})
}
