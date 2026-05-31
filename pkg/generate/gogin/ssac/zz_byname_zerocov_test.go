//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package ssac

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameApplyOperation_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("GetWidget")
	doc := docZeroCov("GetWidget")
	pathItem := doc.Paths.Find("/widgets/{id}")
	v := verbOp{method: "GET", op: pathItem.Get}
	g.applyOperation(pathItem, v, "GetWidget")
	if g.Method != "GET" {
		t.Errorf("applyOperation Method = %q", g.Method)
	}
	if !g.PathParams["id"] {
		t.Errorf("applyOperation did not load path params: %v", g.PathParams)
	}
}

func TestByNameGenerateHTTPMethod_ZeroCov(t *testing.T) {
	doc := docZeroCov("GetWidget")
	sf := ssacparser.ServiceFunc{
		Name:     "GetWidget",
		FileName: "get_widget.ssac",
		Sequences: []ssacparser.Sequence{
			{Type: "response", Inputs: map[string]string{"name": "\"ok\""}},
		},
	}
	fs := &yongol.Fullstack{OpenAPIDoc: doc}
	if err := generateHTTPMethod(sf, fs, t.TempDir(), "example.com/app"); err != nil {
		t.Fatalf("generateHTTPMethod: %v", err)
	}
}

func TestByNameBuildQueryParam_ZeroCov(t *testing.T) {
	// nil ParameterRef → default string.
	if qp := buildQueryParam(nil, "ListItems"); qp.GoType != "string" {
		t.Errorf("nil param GoType = %q", qp.GoType)
	}
	// integer + format + enum.
	sch := openapi3.NewIntegerSchema()
	sch.Format = "int64"
	sch.Enum = []interface{}{int64(1), int64(2)}
	p := &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name:     "status",
		Required: true,
		Schema:   &openapi3.SchemaRef{Value: sch},
	}}
	qp := buildQueryParam(p, "ListItems")
	if !qp.IsRequired || !qp.IsEnum || qp.EnumTypeName == "" {
		t.Errorf("buildQueryParam enum = %+v", qp)
	}
	// schema nil branch.
	p2 := &openapi3.ParameterRef{Value: &openapi3.Parameter{Name: "x"}}
	if qp := buildQueryParam(p2, "Op"); qp.GoType != "string" {
		t.Errorf("no-schema GoType = %q", qp.GoType)
	}
}

func TestByNameBuildResponseConvert_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "GetWorkflow", SuccessStatus: 200}
	// scalar model.
	lines := g.buildResponseConvert("Workflow", "wf")
	if len(lines) != 3 || !strings.Contains(lines[0], "convertWorkflow(wf)") {
		t.Errorf("scalar convert = %v", lines)
	}
	// list model.
	listLines := g.buildResponseConvert("[]Workflow", "wfs")
	if !strings.Contains(listLines[0], "convertWorkflowList(wfs)") {
		t.Errorf("list convert = %v", listLines)
	}
}

func TestByNameRenderPgtypeField_ZeroCov(t *testing.T) {
	g := &methodGen{RespFields: map[string]responseField{
		"name": {JSONName: "name", IsRequired: true},
	}}
	// required → direct.
	if got := g.renderPgtypeField("name", "row.Name", "conv"); !strings.Contains(got, "Name: conv,") {
		t.Errorf("required pgtype field = %q", got)
	}
	// optional + not-already-pointer → ptrOf wrap.
	if got := g.renderPgtypeField("note", "row.Note", "conv"); !strings.Contains(got, "ptrOf(conv)") {
		t.Errorf("optional pgtype field = %q", got)
	}
}

func TestByNameRenderRefResponseField_ZeroCov(t *testing.T) {
	scalarLocal := map[string]string{"workflow": "wfLocal"}
	listLocal := map[string]string{"actions": "actsLocal"}

	// scalar required.
	rf := responseField{JSONName: "workflow", RefType: "Workflow", IsRequired: true}
	got := renderRefResponseField("Workflow", "workflow", rf, scalarLocal, listLocal)
	if !strings.Contains(got, "Workflow: *wfLocal,") {
		t.Errorf("scalar required ref = %q", got)
	}
	// scalar optional.
	rfOpt := responseField{JSONName: "workflow", RefType: "Workflow"}
	gotOpt := renderRefScalarResponseField("Workflow", "workflow", rfOpt, scalarLocal)
	if !strings.Contains(gotOpt, "Workflow: wfLocal,") {
		t.Errorf("scalar optional ref = %q", gotOpt)
	}
	// array branch.
	rfArr := responseField{JSONName: "actions", RefType: "Action", IsArray: true}
	gotArr := renderRefResponseField("Actions", "actions", rfArr, scalarLocal, listLocal)
	if gotArr == "" {
		t.Errorf("array ref empty")
	}
}

func TestByNameRenderResponseFieldHoisted_ZeroCov(t *testing.T) {
	g := &methodGen{RespFields: map[string]responseField{
		"workflow": {JSONName: "workflow", RefType: "Workflow"},
		"count":    {JSONName: "count", IsRequired: true},
	}}
	scalarLocal := map[string]string{"workflow": "wfLocal"}
	listLocal := map[string]string{}

	// $ref → delegate.
	if got := g.renderResponseFieldHoisted("workflow", "x", scalarLocal, listLocal); got == "" {
		t.Errorf("ref hoisted empty")
	}
	// required non-ref → direct.
	if got := g.renderResponseFieldHoisted("count", "n", scalarLocal, listLocal); !strings.Contains(got, "Count: n,") {
		t.Errorf("required hoisted = %q", got)
	}
	// integer literal → int64 wrap + ptrOf.
	if got := g.renderResponseFieldHoisted("total", "5", scalarLocal, listLocal); !strings.Contains(got, "int64(5)") {
		t.Errorf("int literal hoisted = %q", got)
	}
	// variable → &expr.
	if got := g.renderResponseFieldHoisted("name", "v", scalarLocal, listLocal); !strings.Contains(got, "&v,") {
		t.Errorf("var hoisted = %q", got)
	}
}

func TestByNameWrapAuthClaimsFields_ZeroCov(t *testing.T) {
	if got := wrapAuthClaimsFields("auth", "IssueToken", "UserID: id"); !strings.Contains(got, "model.UserClaim{") {
		t.Errorf("issue token wrap = %q", got)
	}
	if got := wrapAuthClaimsFields("auth", "RefreshToken", "UserID: id"); !strings.Contains(got, "model.UserClaim{") {
		t.Errorf("refresh token wrap = %q", got)
	}
	// non-auth passthrough.
	if got := wrapAuthClaimsFields("mail", "Send", "To: x"); got != "To: x" {
		t.Errorf("passthrough = %q", got)
	}
}

func TestByNameWrapInsertExpr_ZeroCov(t *testing.T) {
	g := &methodGen{}
	// alreadyPgtype → passthrough.
	if got, imp := g.wrapInsertExpr("k", "row.Field", true, "request.field"); got != "row.Field" || imp != nil {
		t.Errorf("alreadyPgtype = %q %v", got, imp)
	}
	// no SQLc column → passthrough (lookupSQLCMethodColumn nil).
	if got, _ := g.wrapInsertExpr("unknown", "v", false, "v"); got != "v" {
		t.Errorf("no-col passthrough = %q", got)
	}
}
