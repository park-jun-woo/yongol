//ff:func feature=stml-gen type=test control=sequence
//ff:what TestNestedFetchBindScope — 중첩 fetch의 동명 필드가 각자 op 타입으로 방출되는지(bindCtx.opID 재설정) 검증

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNestedFetchBindScope(t *testing.T) {
	page, _ := stmlparser.ParseReader("nested.html", strings.NewReader(`<main>
  <section data-fetch="OpA">
    <span data-bind="flag"></span>
    <div data-fetch="OpB">
      <span data-bind="flag"></span>
    </div>
  </section>
</main>`))
	opt := GenerateOptions{
		ResponseBindTypes: map[string]map[string]oapiparser.FieldTypeInfo{
			"OpA": {"flag": {Type: "boolean"}},
			"OpB": {"flag": {Type: "string"}},
		},
	}
	code := GeneratePage(page, "", opt)

	// OpA's flag (boolean) → Yes/No against opAData
	assertContains(t, code, "<span>{opAData.flag ? 'Yes' : 'No'}</span>")
	// OpB's flag (string) → plain value against opBData — NOT re-interpreted
	// as boolean from the outer op's scope.
	assertContains(t, code, "<span>{opBData.flag}</span>")
	assertNotContains(t, code, "{opBData.flag ? 'Yes' : 'No'}")
}
