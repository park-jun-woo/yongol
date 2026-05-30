//ff:func feature=validate type=test topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_ZeroCov(t *testing.T) {
	// Empty Fullstack runs every sub-rule without panicking; diags empty.
	diags := Run(&yongol.Fullstack{})
	if len(diags) != 0 {
		t.Fatalf("empty fullstack should yield 0 diags, got %d: %+v", len(diags), diags)
	}
}

func TestCollectInputKeys_ZeroCov(t *testing.T) {
	seq := ssac.Sequence{
		Args: []ssac.Arg{
			{Field: "OrgID"},
			{Field: ""}, // skipped
			{Field: "Page"},
		},
		Inputs: map[string]string{
			"status": "x.Status",
			"":       "ignored", // empty key skipped
		},
	}
	keys := collectInputKeys(seq)
	got := map[string]bool{}
	for _, k := range keys {
		got[k] = true
	}
	for _, want := range []string{"OrgID", "Page", "status"} {
		if !got[want] {
			t.Errorf("missing key %q in %v", want, keys)
		}
	}
	if got[""] {
		t.Error("empty key should be skipped")
	}
}

func TestBuildQueryParamMap_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "UserFindByEmail", Params: []string{"Email"}},
			{Name: "ListAll", Params: nil}, // skipped (no params)
		},
	}
	m := buildQueryParamMap(fs)
	if _, ok := m["ListAll"]; ok {
		t.Error("query with no params should be skipped")
	}
	set := m["UserFindByEmail"]
	if !set["Email"] {
		t.Errorf("expected Email param, got %v", set)
	}
}

func TestBuildXqs18OperationMap_ZeroCov(t *testing.T) {
	// nil doc → empty map.
	if m := buildXqs18OperationMap(nil); len(m) != 0 {
		t.Errorf("nil doc should give empty map, got %v", m)
	}

	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/users", &openapi3.PathItem{
		Get:  &openapi3.Operation{OperationID: "ListUsers"},
		Post: &openapi3.Operation{OperationID: ""}, // skipped: no opID
	})
	m := buildXqs18OperationMap(doc)
	if m["ListUsers"] == nil {
		t.Errorf("expected ListUsers in map, got %v", m)
	}
	if len(m) != 1 {
		t.Errorf("expected exactly 1 op, got %d", len(m))
	}
}

func TestBuildXqs18DDLColumnTypeMap_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"id":   {Name: "id", RawType: "BIGINT", NotNull: true},
				"name": {Name: "name", RawType: "TEXT"},
			},
		}},
	}
	m := buildXqs18DDLColumnTypeMap(fs)
	cols := m["users"]
	if cols == nil {
		t.Fatalf("expected users table, got %v", m)
	}
	if cols["id"] == "" || cols["name"] == "" {
		t.Errorf("expected Go types for columns, got %v", cols)
	}
}

func TestBuildXqs18OAPIParamTypeMap_ZeroCov(t *testing.T) {
	intType := &openapi3.Types{"integer"}
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name: "page",
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type:   intType,
					Format: "int64",
				}},
			}},
			nil, // skipped: nil ref
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name:   "noschema",
				Schema: nil, // skipped: no schema
			}},
		},
	}
	m := buildXqs18OAPIParamTypeMap(op)
	if m["page"] == "" {
		t.Errorf("expected page type, got %v", m)
	}
	if _, ok := m["noschema"]; ok {
		t.Error("param without schema should be skipped")
	}
}

func TestXqs72CheckSeq_ZeroCov(t *testing.T) {
	fn := ssac.ServiceFunc{}
	empty := map[string]string{}
	pm := map[string]map[string]bool{}
	dct := map[string]map[string]string{}
	qbm := map[string]sqlcparser.QuerySpec{}

	// call type → nil.
	if got := xqs72CheckSeq(fn, ssac.Sequence{Type: "call", Model: "X.Y"}, empty, pm, dct, qbm); got != nil {
		t.Error("call type should return nil")
	}
	// empty model → nil.
	if got := xqs72CheckSeq(fn, ssac.Sequence{Type: "get"}, empty, pm, dct, qbm); got != nil {
		t.Error("empty model should return nil")
	}
	// unknown query → nil.
	if got := xqs72CheckSeq(fn, ssac.Sequence{Type: "get", Model: "User.Find"}, empty, pm, dct, qbm); got != nil {
		t.Error("unknown query should return nil")
	}
}

func TestCheckSingleInputKeyCase_ZeroCov(t *testing.T) {
	fn := ssac.ServiceFunc{FileName: "svc.ssac"}
	seq := ssac.Sequence{Line: 5}

	// exact match → no diag.
	if _, fired := checkSingleInputKeyCase(fn, seq, "Email", map[string]bool{"Email": true}); fired {
		t.Error("exact match should not fire")
	}
	// snake match → no diag.
	if _, fired := checkSingleInputKeyCase(fn, seq, "BidAmount", map[string]bool{"bid_amount": true}); fired {
		t.Error("snake-equivalent should not fire")
	}
	// no match at all → no diag.
	if _, fired := checkSingleInputKeyCase(fn, seq, "Email", map[string]bool{"other": true}); fired {
		t.Error("no match should not fire")
	}
	// case-insensitive only match → fires.
	d, fired := checkSingleInputKeyCase(fn, seq, "EMAIL", map[string]bool{"Email": true})
	if !fired {
		t.Fatal("case-insensitive mismatch should fire")
	}
	if d.File != "svc.ssac" || d.Line != 5 {
		t.Errorf("diag fields = %+v", d)
	}
}

func TestCheckSeqInputKeyCase_ZeroCov(t *testing.T) {
	g := &rule.Ground{Lookup: map[string]rule.StringSet{
		"SQLc.param.User": {"Email": true},
	}}

	// call type → nil (early return).
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Type: "call", Model: "User.X"}, g); got != nil {
		t.Error("call type should return nil")
	}
	// package set → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Package: "session", Model: "User.X"}, g); got != nil {
		t.Error("package set should return nil")
	}
	// empty model → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{}, g); got != nil {
		t.Error("empty model should return nil")
	}
	// model without dot → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Model: "User"}, g); got != nil {
		t.Error("model without method should return nil")
	}
	// unknown model params → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Model: "Other.Find", Args: []ssac.Arg{{Field: "Email"}}}, g); got != nil {
		t.Error("unknown model should return nil")
	}
	// no input keys → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Model: "User.Find"}, g); got != nil {
		t.Error("no input keys should return nil")
	}
	// casing mismatch → diag.
	diags := checkSeqInputKeyCase(
		ssac.ServiceFunc{FileName: "s.ssac"},
		ssac.Sequence{Model: "User.Find", Line: 9, Args: []ssac.Arg{{Field: "email"}}},
		g,
	)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
}
