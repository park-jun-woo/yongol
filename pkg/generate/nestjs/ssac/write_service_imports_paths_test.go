//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteServiceImportsPaths — import 경로 ../ 사용 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteServiceImportsPaths(t *testing.T) {
	t.Run("NoOldPaths", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "ListWorkflows",
			HTTPMethod:  "GET",
			Ops: []ir.Op{
				{Kind: ir.OpAuth, Auth: &ir.AuthOp{Action: "list", Resource: "wf"}},
				{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "wf.listed"}},
			},
		}
		var b strings.Builder
		writeServiceImports(&b, plan)
		got := b.String()
		if strings.Contains(got, "../../") {
			t.Errorf("should not contain ../../, got: %s", got)
		}
		if !strings.Contains(got, "../prisma/prisma.service") {
			t.Errorf("expected ../prisma, got: %s", got)
		}
		if !strings.Contains(got, "../queue/queue.service") {
			t.Errorf("expected ../queue, got: %s", got)
		}
		if !strings.Contains(got, "../authz/authz.service") {
			t.Errorf("expected ../authz, got: %s", got)
		}
	})
}
