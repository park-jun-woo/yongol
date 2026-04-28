//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh09CheckFile — 한 hurl 파일에서 사용되지 않은 capture 를 찾는다

package hurl_openapi

import (
	"os"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh09CheckFile reads file once and emits XOH-09 WARNINGs for every
// capture whose variable name never appears again in the body.
func xoh09CheckFile(file string, entries []hurl.HurlEntry) []diagnostic.Diagnostic {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	text := string(content)
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		diags = append(diags, xoh09CheckEntryCaptures(file, text, e)...)
	}
	return diags
}
