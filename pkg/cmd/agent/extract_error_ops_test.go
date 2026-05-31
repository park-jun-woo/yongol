//ff:func feature=agent type=test control=sequence
//ff:what TestExtractErrorOps — nil 에러 처리 및 line/path 매칭으로 원인 op 수집 검증
package agent

import (
	"errors"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestExtractErrorOps(t *testing.T) {
	// nil error returns nothing.
	if ops, rl := extractErrorOps(nil, nil, nil, ""); ops != nil || rl != nil {
		t.Errorf("nil error = %v, %v, want nil, nil", ops, rl)
	}

	offsets := []pathOffset{
		{Op: "OpA", Path: "/users/{id}", StartLine: 1, EndLine: 5},
		{Op: "OpB", Path: "/orders", StartLine: 6, EndLine: 10},
	}

	// Path match collects OpB.
	ops, _ := extractErrorOps(errors.New("validation failed at /orders"), offsets, nil, "")
	if len(ops) != 1 || ops[0] != "OpB" {
		t.Errorf("path match = %v, want [OpB]", ops)
	}

	// Line match (line 3 falls inside OpA's range) collects OpA.
	ops, _ = extractErrorOps(errors.New("yaml: line 3: bad value"), offsets, nil, "")
	if len(ops) != 1 || ops[0] != "OpA" {
		t.Errorf("line match = %v, want [OpA]", ops)
	}

	// Schema match: a $ref schema name maps (case-insensitively) to a feature Op.
	feats := []features.Feature{{Op: "OpA"}}
	ops, _ = extractErrorOps(errors.New("invalid $ref #/components/schemas/opa here"), offsets, feats, "")
	if len(ops) != 1 || ops[0] != "OpA" {
		t.Errorf("schema match = %v, want [OpA]", ops)
	}

	// Grep match: a quoted keyword found in yamlContent at a line within OpB's
	// range collects OpB and populates relativeLines.
	yaml := "line1\nline2\nline3\nline4\nline5\nspecial_field: x\nl7\nl8\nl9\nl10\n"
	ops, rl := extractErrorOps(errors.New(`unknown field "special_field"`), offsets, nil, yaml)
	if len(ops) != 1 || ops[0] != "OpB" {
		t.Errorf("grep match = %v, want [OpB]", ops)
	}
	if rl == nil {
		t.Fatal("expected non-nil relativeLines from grep match")
	}
	if _, ok := rl["OpB"]; !ok {
		t.Errorf("relativeLines missing OpB: %v", rl)
	}
}
