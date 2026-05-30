//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what zz_zerocov_extract — 0% OpenAPI 추출 헬퍼(applyOperation/extractFromOpenAPI/tryExtractFromPathItem/extractBodyFormats/extractRespFields) 검증

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// newMethodGenZeroCov returns a methodGen with all maps initialised so the
// extract helpers can write into them without panicking.
func newMethodGenZeroCov(name string) *methodGen {
	return &methodGen{
		FuncName:        name,
		SuccessStatus:   200,
		PathParams:      map[string]bool{},
		QueryParams:     map[string]queryParam{},
		BodyFormats:     map[string]string{},
		RespFields:      map[string]responseField{},
		BodyJSONBFields: map[string]bool{},
		DeclaredVars:    map[string]bool{},
	}
}

// docZeroCov builds a minimal OpenAPI doc with one GET operation that has a
// path param, query param, JSON request body (with format + enum props), and
// a 200 response referencing component schemas.
func docZeroCov(operationID string) *openapi3.T {
	strType := &openapi3.Types{"string"}
	reqSchema := openapi3.NewSchema()
	reqSchema.Type = &openapi3.Types{"object"}
	reqSchema.Required = []string{"email"}
	reqSchema.Properties = openapi3.Schemas{
		"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType, Format: "email"}},
		"plan":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType, Enum: []any{"free", "pro"}}},
		"meta":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, AdditionalProperties: openapi3.AdditionalProperties{Has: boolPtr(true)}}},
	}

	respSchema := openapi3.NewSchema()
	respSchema.Type = &openapi3.Types{"object"}
	respSchema.Required = []string{"widget"}
	respSchema.Properties = openapi3.Schemas{
		"widget": &openapi3.SchemaRef{Ref: "#/components/schemas/Widget"},
		"name":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType}},
	}

	op := openapi3.NewOperation()
	op.OperationID = operationID
	op.Parameters = openapi3.Parameters{
		{Value: &openapi3.Parameter{Name: "id", In: "path", Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType}}}},
		{Value: &openapi3.Parameter{Name: "q", In: "query", Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType}}}},
	}
	op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchema(reqSchema)}
	resp := openapi3.NewResponse().WithJSONSchema(respSchema)
	op.Responses = openapi3.NewResponses()
	op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})

	pathItem := &openapi3.PathItem{Get: op}
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/widgets/{id}", pathItem)
	return doc
}

func boolPtr(b bool) *bool { return &b }

//ff:what TestExtractFromOpenAPI_ZeroCov — operationId 매칭 후 path/query/body/resp 메타 적재
func TestExtractFromOpenAPI_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("GetWidget")
	doc := docZeroCov("GetWidget")
	g.extractFromOpenAPI(doc, "GetWidget")

	if !g.PathParams["id"] {
		t.Errorf("expected path param id, got %v", g.PathParams)
	}
	if _, ok := g.QueryParams["q"]; !ok {
		t.Errorf("expected query param q, got %v", g.QueryParams)
	}
	if g.BodyFormats["email"] != "email" {
		t.Errorf("expected email format, got %v", g.BodyFormats)
	}
	if g.BodyFormats["plan"] != "enum" {
		t.Errorf("expected plan enum, got %v", g.BodyFormats)
	}
	if !g.BodyJSONBFields["meta"] {
		t.Errorf("expected meta marked JSONB, got %v", g.BodyJSONBFields)
	}
	if !g.BodyRequiredFields["email"] {
		t.Errorf("expected email required, got %v", g.BodyRequiredFields)
	}
	wf, ok := g.RespFields["widget"]
	if !ok || wf.RefType != "Widget" {
		t.Errorf("expected widget RefType Widget, got %+v (ok=%v)", wf, ok)
	}
	if g.Method != "GET" {
		t.Errorf("Method = %q, want GET", g.Method)
	}
}

//ff:what TestExtractFromOpenAPI_NoMatch_ZeroCov — operationId 미매칭 / Paths nil
func TestExtractFromOpenAPI_NoMatch_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("Nope")
	doc := docZeroCov("GetWidget")
	g.extractFromOpenAPI(doc, "DoesNotExist")
	if len(g.PathParams) != 0 {
		t.Errorf("expected no params on mismatch, got %v", g.PathParams)
	}

	// Nil paths — early return.
	g2 := newMethodGenZeroCov("X")
	g2.extractFromOpenAPI(&openapi3.T{}, "anything")
	if len(g2.PathParams) != 0 {
		t.Errorf("expected no-op for nil paths")
	}
}

//ff:what TestExtractBodyFormats_NoBody_ZeroCov — request body 없음 early-return
func TestExtractBodyFormats_NoBody_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("X")
	op := openapi3.NewOperation()
	g.extractBodyFormats(op) // op.RequestBody == nil → returns
	if len(g.BodyFormats) != 0 {
		t.Errorf("expected empty BodyFormats, got %v", g.BodyFormats)
	}
}

//ff:what TestExtractRespFields_NoResp_ZeroCov — 매칭 응답 없음 early-return
func TestExtractRespFields_NoResp_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("X")
	op := openapi3.NewOperation()
	op.Responses = openapi3.NewResponses()
	g.extractRespFields(op) // no 200 → returns
	if len(g.RespFields) != 0 {
		t.Errorf("expected empty RespFields, got %v", g.RespFields)
	}
}

//ff:what TestNewMethodGen_ZeroCov — newMethodGen 생성 + OpenAPI 메타 주입 + VarTypes/ImportMap 사전계산
func TestNewMethodGen_ZeroCov(t *testing.T) {
	doc := docZeroCov("GetWidget")
	sf := ssacparser.ServiceFunc{
		Name:     "GetWidget",
		FileName: "get_widget.ssac",
		Imports:  []string{"example.com/app/internal/dashboard"},
		Sequences: []ssacparser.Sequence{
			{Type: "get", Model: "Widget.FindByID", Result: &ssacparser.Result{Type: "Widget", Var: "widget"}},
		},
	}
	g := newMethodGen(doc, sf, "example.com/app", false, nil, nil, false, nil, nil, nil, nil, nil)
	if g == nil {
		t.Fatal("newMethodGen returned nil")
	}
	if g.FuncName != "GetWidget" {
		t.Errorf("FuncName = %q", g.FuncName)
	}
	// VarTypes precomputed from the get sequence's Result binding.
	if g.VarTypes["widget"] != "Widget" {
		t.Errorf("VarTypes = %v, want widget→Widget", g.VarTypes)
	}
	// ImportMap keyed by path.Base.
	if g.ImportMap["dashboard"] != "example.com/app/internal/dashboard" {
		t.Errorf("ImportMap = %v", g.ImportMap)
	}
	// OpenAPI metadata injected via extractFromOpenAPI (path param "id").
	if !g.PathParams["id"] {
		t.Errorf("expected path param id injected, got %v", g.PathParams)
	}

	// useTx=true path: FirstErr=false, queryVar=qtx.
	g2 := newMethodGen(doc, sf, "m", true, nil, nil, true, nil, nil, nil, nil, nil)
	if g2.FirstErr {
		t.Errorf("useTx → FirstErr should be false")
	}
	if g2.queryVar() != "qtx" {
		t.Errorf("useTx → queryVar should be qtx, got %q", g2.queryVar())
	}
}

//ff:what TestTryExtractFromPathItem_NoVerbMatch_ZeroCov — verb 미매칭 시 false
func TestTryExtractFromPathItem_NoVerbMatch_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("X")
	doc := docZeroCov("GetWidget")
	pathItem := doc.Paths.Find("/widgets/{id}")
	if g.tryExtractFromPathItem(pathItem, "OtherOp") {
		t.Errorf("expected false for non-matching operationId")
	}
}
