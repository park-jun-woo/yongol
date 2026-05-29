//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestExtractErrorOps — nil 에러 처리 및 line/path 매칭으로 원인 op 수집 검증

package agent

import (
	"errors"
	"testing"
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
}
