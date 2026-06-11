//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what xoe02Detail — 표시 필드 부재(스키마 라인) vs 비-string(프로퍼티 라인) 메시지/라인 분기 검증
package openapi

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestXOE02Detail(t *testing.T) {
	lines := &oapiparser.LineIndex{
		Schemas:          map[string]int{"ErrorResponse": 10},
		SchemaProperties: map[string]map[string]int{"ErrorResponse": {"error": 11}},
	}

	// no candidate property → schema line + "없어" message
	line, msg := xoe02Detail("ErrorResponse", false, false, lines)
	if line != 10 || !strings.Contains(msg, "없어") {
		t.Errorf("missing-field detail: line=%d msg=%q", line, msg)
	}

	// non-string error present → property line + "string 타입이 아니" message
	line, msg = xoe02Detail("ErrorResponse", true, false, lines)
	if line != 11 || !strings.Contains(msg, "string 타입이 아니") {
		t.Errorf("non-string-error detail: line=%d msg=%q", line, msg)
	}

	// only message present → field "message" named in message
	_, msg = xoe02Detail("ErrorResponse", false, true, lines)
	if !strings.Contains(msg, "`message`") {
		t.Errorf("message-field detail: msg=%q", msg)
	}
}
