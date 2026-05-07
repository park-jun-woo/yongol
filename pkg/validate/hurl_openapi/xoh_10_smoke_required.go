//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-10 — smoke.hurl 이 specs/tests/ 에 존재해야 함

package hurl_openapi

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xoh10SmokeRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	for _, f := range fs.HurlFiles {
		if filepath.Base(f) == "smoke.hurl" {
			return nil
		}
	}
	return []diagnostic.Diagnostic{{
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XOH-10] smoke.hurl is required but not found in specs/tests/",
		Advice:  "Create specs/tests/smoke.hurl with at least one request per endpoint",
	}}
}
