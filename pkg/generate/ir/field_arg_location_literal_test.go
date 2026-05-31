//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgLocation -- FieldArg.Location 분류 검증 (path/query/body/var/literal/user)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgLocationLiteral(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqPut,
				Model: "Workflow.UpdateStatus",
				Inputs: map[string]string{
					"Status": `"archived"`,
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	putOp := plan.Ops[0]
	statusArg := findArgByKey(putOp.Put.Args, "Status")
	if statusArg == nil {
		t.Fatal("missing Status arg")
	}
	if statusArg.Location != LocLiteral {
		t.Errorf("Status.Location = %q, want %q", statusArg.Location, LocLiteral)
	}
}
