//ff:func feature=crosscheck type=test control=sequence
//ff:what TestByName_ZeroCov — hurl 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package hurl

import (
	"strings"
	"testing"
)

func TestByNameReplaceHurlVarAt_ZeroCov(t *testing.T) {
	// not a var start.
	var b1 strings.Builder
	if _, ok := replaceHurlVarAt("abc", 0, &b1); ok {
		t.Errorf("replaceHurlVarAt non-var should be false")
	}

	// quoted var -> bare.
	var b2 strings.Builder
	b2.WriteByte('"')
	body := `"{{token}}"`
	if _, ok := replaceHurlVarAt(body, 1, &b2); !ok {
		t.Errorf("replaceHurlVarAt quoted var should be true")
	}

	// unquoted var.
	var b3 strings.Builder
	if _, ok := replaceHurlVarAt(`{{token}}`, 0, &b3); !ok {
		t.Errorf("replaceHurlVarAt unquoted var should be true")
	}

	// unterminated var.
	var b4 strings.Builder
	if _, ok := replaceHurlVarAt(`{{token`, 0, &b4); !ok {
		t.Errorf("replaceHurlVarAt unterminated should be true")
	}
}
