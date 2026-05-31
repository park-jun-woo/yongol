//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildPutAndDelete_ZeroCov — @put / @delete (둘 다 buildPut)
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
