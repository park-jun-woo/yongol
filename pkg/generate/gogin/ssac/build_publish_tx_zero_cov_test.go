//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildPublishTx_ZeroCov — UseTx publish + subscribe err 전파 분기
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
