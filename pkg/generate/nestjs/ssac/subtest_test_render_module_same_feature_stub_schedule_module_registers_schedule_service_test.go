//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what subtestTestRenderModuleSameFeatureStubScheduleModuleRegistersScheduleService — ScheduleModuleRegistersScheduleService 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestRenderModuleSameFeatureStubScheduleModuleRegistersScheduleService(t *testing.T) {

	plans := []*ir.ServicePlan{
		{
			OperationID: "RunSchedule",
			HTTPMethod:  "POST",
			Feature:     "schedule",
			Ops: []ir.Op{
				{Kind: ir.OpCall, Call: &ir.CallOp{
					Package:       "schedule",
					TargetFeature: "schedule",
					Function:      "ParseCron",
				}},
			},
		},
	}
	got, err := RenderModule("schedule", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	if !strings.Contains(got, "ScheduleService,") {
		t.Errorf("expected ScheduleService in providers, got:\n%s", got)
	}
	if !strings.Contains(got, "import { ScheduleService } from './schedule.service'") {
		t.Errorf("expected ScheduleService import, got:\n%s", got)
	}

}
