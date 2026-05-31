//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what zz_zerocov_test — fastapi.buildPlansByFeature 0% 커버리지 단위 테스트
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildPlansByFeature_ZeroCov(t *testing.T) {
	// Empty fullstack → empty map (function entry + loop skipped).
	if m := buildPlansByFeature(&yongol.Fullstack{}); len(m) != 0 {
		t.Errorf("empty fullstack should yield empty map, got %v", m)
	}

	// One auth-style service func → builds a plan grouped by feature.
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "ArchiveWorkflow",
			FileName: "archive_workflow.ssac",
			Feature:  "workflow",
			Sequences: []ssac.Sequence{{
				Type:      ssac.SeqAuth,
				Action:    "ArchiveWorkflow",
				Resource:  "workflow",
				Inputs:    map[string]string{"ResourceID": "wf.ID"},
				Message:   "Forbidden",
				ErrStatus: 403,
			}},
		}},
	}
	m := buildPlansByFeature(fs)
	total := 0
	for _, plans := range m {
		total += len(plans)
	}
	if total != 1 {
		t.Fatalf("expected 1 plan total, got %d (map=%v)", total, m)
	}
}
