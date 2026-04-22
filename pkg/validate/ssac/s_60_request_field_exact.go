//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-60 — SSaC 에서 사용한 request.<field> 가 OpenAPI schema 에 case-exact 로 존재

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s60RequestFieldExact validates S-60: every `request.<field>` referenced by
// SSaC must appear case-exactly in the OpenAPI operation's request schema.
// Manual states: "request.* field names must exactly match the OpenAPI
// request schema property names." — violation is ERROR (codegen will produce
// non-compilable bindings).
//
// Complementary to S-51 (OpenAPI→SSaC direction, WARNING for unused).
func s60RequestFieldExact(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		// request.* 은 requestBody 필드뿐 아니라 path/query parameter 도 포함
		expected := map[string]bool{}
		if body := g.Lookup["OpenAPI.request."+fn.Name]; body != nil {
			for k := range body {
				expected[k] = true
			}
		}
		if params := g.Lookup["OpenAPI.param."+fn.Name]; params != nil {
			for k := range params {
				expected[k] = true
			}
		}
		if len(expected) == 0 {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "request" && arg.Field != "" && !expected[arg.Field] {
					diags = append(diags, missingRequestField(fn, seq, arg.Field))
				}
			}
			for _, val := range seq.Inputs {
				v := strings.TrimSpace(val)
				if !strings.HasPrefix(v, "request.") {
					continue
				}
				field := strings.TrimPrefix(v, "request.")
				if field == "" {
					continue
				}
				if !expected[field] {
					diags = append(diags, missingRequestField(fn, seq, field))
				}
			}
		}
	}
	return diags
}
