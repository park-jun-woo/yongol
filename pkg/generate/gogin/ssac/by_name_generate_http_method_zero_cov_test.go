//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
	if err := generateHTTPMethod(sf, fs, t.TempDir(), "example.com/app", nil); err != nil {
		t.Fatalf("generateHTTPMethod: %v", err)
	}
}
