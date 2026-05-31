//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchStringLiteralBoundToDateTimeFieldErrorsFalseNegativeClosed — string literal bound to date-time field errors (false negative closed) 서브테스트
package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchStringLiteralBoundToDateTimeFieldErrorsFalseNegativeClosed(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name:     "approveItem",
				FileName: "item.ssac",
				Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"approved_at": `"2026-01-01T00:00:00Z"`}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			// expected is time.Time (format: date-time); a string literal
			// (actual="string") must now ERROR — previously slipped through.
			"OpenAPI.response.approveItem.approved_at": "time.Time",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "XOS-67") {
		t.Errorf("Message missing XOS-67: %s", diags[0].Message)
	}

}
