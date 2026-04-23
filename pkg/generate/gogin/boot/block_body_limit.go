//ff:func feature=gen-gogin type=generator control=sequence topic=dos-guard
//ff:what blockBodyLimit — middleware.BodyLimit / MultipartLimit / OverrideBodyLimit 등록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockBodyLimit emits the body-limit middleware registration. Active
// always (defaults applied when manifest.backend.http is nil). env overrides
// (BACKEND_HTTP_BODY_LIMIT, BACKEND_HTTP_MULTIPART_LIMIT) resolved at
// runtime via envInt64 (emitted into cmd/ helpers by blockEnvHelpers
// extension below).
//
// Order: registered AFTER RequestValidator block reference but must run
// BEFORE it in the gin chain so MaxBytesReader is attached before the
// validator calls io.ReadAll(req.Body). This is why blockBodyLimit is
// placed before blockRequestValidator in collectActiveBlocks.
func blockBodyLimit(fs *yongol.Fullstack, modulePath string) MainBlock {
	bodyLimit, multipartLimit, bodyOverrides, multipartOverrides := resolveHTTPLimits(fs)

	lines := []string{
		fmt.Sprintf(`bodyLimit := envInt64("BACKEND_HTTP_BODY_LIMIT", %d)`, bodyLimit),
		fmt.Sprintf(`multipartLimit := envInt64("BACKEND_HTTP_MULTIPART_LIMIT", %d)`, multipartLimit),
		`r.Use(middleware.BodyLimit(bodyLimit))`,
		`r.Use(middleware.MultipartLimit(multipartLimit))`,
	}

	// Only emit OverrideBodyLimit when there is at least one override.
	if len(bodyOverrides) > 0 || len(multipartOverrides) > 0 {
		lines = append(lines, "bodyOverrides := map[string]int64{")
		for _, k := range sortedKeys(bodyOverrides) {
			lines = append(lines, fmt.Sprintf("\t%q: %d,", k, bodyOverrides[k]))
		}
		lines = append(lines, "}")
		lines = append(lines, "multipartOverrides := map[string]int64{")
		for _, k := range sortedKeys(multipartOverrides) {
			lines = append(lines, fmt.Sprintf("\t%q: %d,", k, multipartOverrides[k]))
		}
		lines = append(lines, "}")
		lines = append(lines, `r.Use(middleware.OverrideBodyLimit(bodyOverrides, multipartOverrides))`)
	}

	return MainBlock{
		Name: "body-limit",
		Imports: []string{
			fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
		},
		Lines: lines,
	}
}
