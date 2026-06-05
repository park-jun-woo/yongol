//ff:func feature=stml-gen type=test control=sequence
//ff:what TestResolveInvalidateOps_Merge — 자동 추론 + 명시 data-invalidates 합집합·중복 제거

package stml

import (
	"reflect"
	"testing"
)

func TestResolveInvalidateOps_Merge(t *testing.T) {
	// No explicit invalidates: output equals auto-inferred set unchanged.
	auto := resolveInvalidateOps("CreateWorkflow", []string{"ListWorkflows"}, nil, nil)
	if want := []string{"ListWorkflows"}; !reflect.DeepEqual(auto, want) {
		t.Errorf("auto-only = %v, want %v", auto, want)
	}

	// Explicit op merged in (union), auto order preserved first.
	merged := resolveInvalidateOps("CreateWorkflow", []string{"ListWorkflows"}, nil, []string{"GetWorkflow"})
	if want := []string{"ListWorkflows", "GetWorkflow"}; !reflect.DeepEqual(merged, want) {
		t.Errorf("merged = %v, want %v", merged, want)
	}

	// Duplicate between auto and explicit is removed.
	dedup := resolveInvalidateOps("CreateWorkflow", []string{"ListWorkflows"}, nil, []string{"ListWorkflows", "GetWorkflow"})
	if want := []string{"ListWorkflows", "GetWorkflow"}; !reflect.DeepEqual(dedup, want) {
		t.Errorf("dedup = %v, want %v", dedup, want)
	}

	// Scoped auto set (action inside a fetch) merged with explicit op.
	scoped := resolveInvalidateOps("CreateAction", nil, map[string][]string{"CreateAction": {"GetWorkflow"}}, []string{"ListActions"})
	if want := []string{"GetWorkflow", "ListActions"}; !reflect.DeepEqual(scoped, want) {
		t.Errorf("scoped = %v, want %v", scoped, want)
	}
}
