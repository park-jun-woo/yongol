//ff:func feature=stml-gen type=test control=sequence
//ff:what TestResolveInvalidateOps_Merge — 자동 추론 + 명시 data-invalidates 합집합·중복 제거, delete 자기 GET removeQueries 분리

package stml

import (
	"reflect"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestResolveInvalidateOps_Merge(t *testing.T) {
	// No explicit invalidates: output equals auto-inferred set unchanged.
	inv, rem := resolveInvalidateOps(stmlparser.ActionBlock{OperationID: "CreateWorkflow"}, []string{"ListWorkflows"}, nil, nil)
	if want := []string{"ListWorkflows"}; !reflect.DeepEqual(inv, want) {
		t.Errorf("auto-only = %v, want %v", inv, want)
	}
	if rem != nil {
		t.Errorf("auto-only remove = %v, want nil", rem)
	}

	// Explicit op merged in (union), auto order preserved first.
	merged, _ := resolveInvalidateOps(stmlparser.ActionBlock{OperationID: "CreateWorkflow", Invalidates: []string{"GetWorkflow"}}, []string{"ListWorkflows"}, nil, nil)
	if want := []string{"ListWorkflows", "GetWorkflow"}; !reflect.DeepEqual(merged, want) {
		t.Errorf("merged = %v, want %v", merged, want)
	}

	// Duplicate between auto and explicit is removed.
	dedup, _ := resolveInvalidateOps(stmlparser.ActionBlock{OperationID: "CreateWorkflow", Invalidates: []string{"ListWorkflows", "GetWorkflow"}}, []string{"ListWorkflows"}, nil, nil)
	if want := []string{"ListWorkflows", "GetWorkflow"}; !reflect.DeepEqual(dedup, want) {
		t.Errorf("dedup = %v, want %v", dedup, want)
	}

	// Scoped auto set (action inside a fetch) merged with explicit op.
	scoped, _ := resolveInvalidateOps(stmlparser.ActionBlock{OperationID: "CreateAction", Invalidates: []string{"ListActions"}}, nil, map[string][]string{"CreateAction": {"GetWorkflow"}}, nil)
	if want := []string{"GetWorkflow", "ListActions"}; !reflect.DeepEqual(scoped, want) {
		t.Errorf("scoped = %v, want %v", scoped, want)
	}

	// Delete action: the self GET (same path-param signature) is split out of
	// invalidate into remove; sibling queries stay invalidated (BUG-132 132-2).
	ppt := map[string]map[string]string{
		"DeleteBuilding": {"buildingId": "integer"},
		"GetBuilding":    {"buildingId": "integer"},
	}
	delInv, delRem := resolveInvalidateOps(
		stmlparser.ActionBlock{OperationID: "DeleteBuilding"},
		[]string{"GetBuilding", "ListBuildingPhotos"}, nil, ppt)
	if want := []string{"ListBuildingPhotos"}; !reflect.DeepEqual(delInv, want) {
		t.Errorf("delete invalidate = %v, want %v", delInv, want)
	}
	if want := []string{"GetBuilding"}; !reflect.DeepEqual(delRem, want) {
		t.Errorf("delete remove = %v, want %v", delRem, want)
	}

	// Delete with no matching self GET on the page: nothing removed.
	noSelfInv, noSelfRem := resolveInvalidateOps(
		stmlparser.ActionBlock{OperationID: "DeleteBuilding"},
		[]string{"ListBuildingPhotos"}, nil, ppt)
	if want := []string{"ListBuildingPhotos"}; !reflect.DeepEqual(noSelfInv, want) {
		t.Errorf("delete no-self invalidate = %v, want %v", noSelfInv, want)
	}
	if noSelfRem != nil {
		t.Errorf("delete no-self remove = %v, want nil", noSelfRem)
	}
}
