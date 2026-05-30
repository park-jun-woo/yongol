//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what zz_zerocov_seq — 0% 커버리지 시퀀스 빌더(@get/@post/@put/@delete/@empty/@exists/@state/publishTx) 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

//ff:what TestBuildGet_ZeroCov — @get + 후속 @empty/@exists 관용 분기
func TestBuildGet_ZeroCov(t *testing.T) {
	g := &methodGen{
		FuncName:     "GetWidget",
		BodyFormats:  map[string]string{},
		RespFields:   map[string]responseField{},
		DeclaredVars: map[string]bool{},
	}
	seq := ssacparser.Sequence{
		Type:   "get",
		Model:  "Widget.FindByID",
		Result: &ssacparser.Result{Type: "Widget", Var: "widget"},
		Inputs: map[string]string{},
		Args:   []ssacparser.Arg{{Source: "request", Field: "ID"}},
	}

	// No following guard → plain err handler.
	lines, _ := g.buildGet(seq, nil)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "server.Queries.WidgetFindByID(") {
		t.Fatalf("expected sqlc call, got:\n%s", body)
	}
	if strings.Contains(body, "pgx.ErrNoRows") {
		t.Fatalf("did not expect ErrNoRows tolerance without guard, got:\n%s", body)
	}

	// Following @empty targeting the same var → ErrNoRows tolerance.
	g.FirstErr = false
	next := &ssacparser.Sequence{Type: "empty", Target: "widget"}
	lines2, imports2 := g.buildGet(seq, next)
	body2 := strings.Join(lines2, "\n")
	if !strings.Contains(body2, "pgx.ErrNoRows") {
		t.Fatalf("expected ErrNoRows tolerance with @empty guard, got:\n%s", body2)
	}
	if !strings.Contains(strings.Join(imports2, " "), "pgx") {
		t.Fatalf("expected pgx import, got %v", imports2)
	}
}

//ff:what TestBuildPost_ZeroCov — @post 위임 (buildGet, next=nil)
func TestBuildPost_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "CreateWidget", BodyFormats: map[string]string{}, DeclaredVars: map[string]bool{}}
	seq := ssacparser.Sequence{
		Type:   "post",
		Model:  "Widget.Create",
		Result: &ssacparser.Result{Type: "Widget", Var: "created"},
		Inputs: map[string]string{},
	}
	lines, _ := g.buildPost(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "WidgetCreate(") {
		t.Fatalf("expected INSERT call, got:\n%s", body)
	}
}

//ff:what TestBuildPutAndDelete_ZeroCov — @put / @delete (둘 다 buildPut)
func TestBuildPutAndDelete_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "UpdateWidget", BodyFormats: map[string]string{}}
	putSeq := ssacparser.Sequence{Type: "put", Model: "Widget.Update", Inputs: map[string]string{}}
	lines, _ := g.buildPut(putSeq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "WidgetUpdate(") {
		t.Fatalf("expected UPDATE call, got:\n%s", body)
	}
	if !strings.Contains(body, "if err != nil") {
		t.Fatalf("expected err check, got:\n%s", body)
	}

	g2 := &methodGen{FuncName: "DeleteWidget", BodyFormats: map[string]string{}}
	delSeq := ssacparser.Sequence{Type: "delete", Model: "Widget.Delete", Inputs: map[string]string{}}
	dlines, _ := g2.buildDelete(delSeq)
	dbody := strings.Join(dlines, "\n")
	if !strings.Contains(dbody, "WidgetDelete(") {
		t.Fatalf("expected DELETE call, got:\n%s", dbody)
	}
}

//ff:what TestBuildEmptyExists_ZeroCov — @empty/@exists guard (col==nil 분기) + subscribe 분기
func TestBuildEmptyExists_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "GetWidget"}
	emptySeq := ssacparser.Sequence{Type: "empty", Target: "widget"}
	lines, imports := g.buildEmpty(emptySeq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "return api.GetWidget404JSONResponse") {
		t.Fatalf("expected 404 guard return, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(imports, " "), "log/slog") {
		t.Fatalf("expected slog import, got %v", imports)
	}

	existsSeq := ssacparser.Sequence{Type: "exists", Target: "widget"}
	elines, _ := g.buildExists(existsSeq)
	ebody := strings.Join(elines, "\n")
	if !strings.Contains(ebody, "return api.GetWidget409JSONResponse") {
		t.Fatalf("expected 409 guard return, got:\n%s", ebody)
	}

	// Subscribe variant adds fmt import + fmt.Errorf return.
	sub := &methodGen{FuncName: "OnWidget", IsSubscribe: true}
	slines, simports := sub.buildEmpty(emptySeq)
	if !strings.Contains(strings.Join(slines, "\n"), "fmt.Errorf") {
		t.Fatalf("expected fmt.Errorf for subscribe, got:\n%s", strings.Join(slines, "\n"))
	}
	if !strings.Contains(strings.Join(simports, " "), `"fmt"`) {
		t.Fatalf("expected fmt import for subscribe, got %v", simports)
	}
}

//ff:what TestBuildState_ZeroCov — @state 전이 검증 (Symbol fallback 포함)
func TestBuildState_ZeroCov(t *testing.T) {
	g := &methodGen{
		FuncName:      "CancelReservation",
		ModulePath:    "example.com/app",
		DiagramSymbol: map[string]string{"reservation": "Reservation"},
	}
	seq := ssacparser.Sequence{
		Type:       "state",
		DiagramID:  "reservation",
		Inputs:     map[string]string{"status": "reservation.Status"},
		Transition: "cancel",
	}
	lines, imports := g.buildState(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "statemachine.ReservationCanTransition(") {
		t.Fatalf("expected statemachine call, got:\n%s", body)
	}
	if !strings.Contains(body, `"cancel"`) {
		t.Fatalf("expected transition literal, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(imports, " "), "internal/statemachine") {
		t.Fatalf("expected statemachine import, got %v", imports)
	}

	// Fallback: DiagramSymbol missing → uses DiagramID.
	g2 := &methodGen{FuncName: "X", ModulePath: "m", DiagramSymbol: map[string]string{}}
	seq2 := ssacparser.Sequence{Type: "state", DiagramID: "order", Inputs: map[string]string{"s": "o.S"}, Transition: "ship"}
	body2 := strings.Join(func() []string { l, _ := g2.buildState(seq2); return l }(), "\n")
	if !strings.Contains(body2, "statemachine.orderCanTransition(") {
		t.Fatalf("expected fallback to DiagramID, got:\n%s", body2)
	}
}

//ff:what TestBuildPublishTx_ZeroCov — UseTx publish + subscribe err 전파 분기
func TestBuildPublishTx_ZeroCov(t *testing.T) {
	seq := ssacparser.Sequence{Type: "publish", Topic: "order.completed"}
	fields := []string{"\t\"id\": order.ID,"}

	lines := buildPublishTx(seq, fields, nil, false)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, `queue.PublishTx(ctx, tx, "order.completed"`) {
		t.Fatalf("expected PublishTx call, got:\n%s", body)
	}
	if !strings.Contains(body, "return nil, err") {
		t.Fatalf("expected handler err return, got:\n%s", body)
	}

	subLines := buildPublishTx(seq, fields, nil, true)
	subBody := strings.Join(subLines, "\n")
	if !strings.Contains(subBody, "return err") || strings.Contains(subBody, "return nil, err") {
		t.Fatalf("expected subscribe single-err return, got:\n%s", subBody)
	}
}

//ff:what TestBuildResponseField_ZeroCov — $ref / array-of-$ref / 원시 / nil 분기
func TestBuildResponseField_ZeroCov(t *testing.T) {
	// nil propRef
	if rf := buildResponseField("Name", nil, true); rf.RefType != "" || !rf.IsRequired {
		t.Fatalf("nil propRef: unexpected %+v", rf)
	}
	// direct $ref
	ref := &openapi3.SchemaRef{Ref: "#/components/schemas/Widget"}
	if rf := buildResponseField("widget", ref, false); rf.RefType != "Widget" || rf.IsArray {
		t.Fatalf("direct ref: unexpected %+v", rf)
	}
	// array of $ref
	arr := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Ref: "#/components/schemas/Item"},
	}}
	if rf := buildResponseField("items", arr, true); rf.RefType != "Item" || !rf.IsArray {
		t.Fatalf("array ref: unexpected %+v", rf)
	}
	// primitive (no ref, not array)
	prim := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if rf := buildResponseField("title", prim, false); rf.RefType != "" || rf.IsArray {
		t.Fatalf("primitive: unexpected %+v", rf)
	}
	// array but items has no ref → no RefType
	arrNoRef := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}}
	if rf := buildResponseField("tags", arrNoRef, false); rf.RefType != "" {
		t.Fatalf("array no-ref items: unexpected %+v", rf)
	}
}
