//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanUnknownType -- 미지원 시퀀스 타입 에러 경로 검증

package ir

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanUnknownType(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name: "Bad",
		Sequences: []ssac.Sequence{
			{Type: "unknown-type"},
		},
	}

	_, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err == nil {
		t.Fatal("expected error for unknown sequence type, got nil")
	}
	if !strings.Contains(err.Error(), "unknown sequence type") {
		t.Errorf("error = %q, want to contain %q", err.Error(), "unknown sequence type")
	}
}
