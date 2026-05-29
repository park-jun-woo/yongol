//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderModuleSameFeatureStub — 같은 feature @call 시 stub 서비스 DI 등록 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderModuleSameFeatureStub(t *testing.T) {
	t.Run("AuthModuleRegistersAuthService", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "Login",
				HTTPMethod:  "POST",
				Feature:     "auth",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{
						Package:       "auth",
						TargetFeature: "auth",
						Function:      "IssueToken",
					}},
				},
			},
		}
		got, err := RenderModule("auth", plans)
		if err != nil {
			t.Fatalf("RenderModule: %v", err)
		}

		// AuthService must be in providers.
		if !strings.Contains(got, "AuthService,") {
			t.Errorf("expected AuthService in providers, got:\n%s", got)
		}

		// AuthService must be in exports.
		exportsIdx := strings.Index(got, "exports: [")
		if exportsIdx < 0 {
			t.Fatal("expected exports array")
		}
		exportsSection := got[exportsIdx:]
		if !strings.Contains(exportsSection, "AuthService,") {
			t.Errorf("expected AuthService in exports, got:\n%s", exportsSection)
		}

		// AuthService import statement.
		if !strings.Contains(got, "import { AuthService } from './auth.service'") {
			t.Errorf("expected AuthService import, got:\n%s", got)
		}
	})

	t.Run("NoCrossFeatureCallNoStub", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateOrder",
				HTTPMethod:  "POST",
				Feature:     "order",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{
						Package:       "billing",
						TargetFeature: "billing",
						Function:      "HoldEscrow",
					}},
				},
			},
		}
		got, err := RenderModule("order", plans)
		if err != nil {
			t.Fatalf("RenderModule: %v", err)
		}

		// OrderService stub import should NOT be present (billing is cross-feature).
		if strings.Contains(got, "import { OrderService }") {
			t.Errorf("should not import OrderService for cross-feature call, got:\n%s", got)
		}
	})

	t.Run("MixedSameAndCrossFeature", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "Login",
				HTTPMethod:  "POST",
				Feature:     "auth",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{
						Package:       "auth",
						TargetFeature: "auth",
						Function:      "IssueToken",
					}},
					{Kind: ir.OpCall, Call: &ir.CallOp{
						Package:       "notification",
						TargetFeature: "notification",
						Function:      "SendWelcome",
					}},
				},
			},
		}
		got, err := RenderModule("auth", plans)
		if err != nil {
			t.Fatalf("RenderModule: %v", err)
		}

		// AuthService stub must be present.
		if !strings.Contains(got, "AuthService,") {
			t.Errorf("expected AuthService in providers, got:\n%s", got)
		}
		// Cross-feature module must be imported.
		if !strings.Contains(got, "NotificationModule") {
			t.Errorf("expected NotificationModule import, got:\n%s", got)
		}
	})

	t.Run("NoSameFeatureCallNoStub", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "ListItems",
				HTTPMethod:  "GET",
				Feature:     "item",
				Ops: []ir.Op{
					{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Item"}},
				},
			},
		}
		got, err := RenderModule("item", plans)
		if err != nil {
			t.Fatalf("RenderModule: %v", err)
		}

		// No stub service should be added.
		if strings.Contains(got, "ItemService,") {
			t.Errorf("should not add ItemService stub without same-feature @call, got:\n%s", got)
		}
	})

	t.Run("ScheduleModuleRegistersScheduleService", func(t *testing.T) {
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
	})
}
