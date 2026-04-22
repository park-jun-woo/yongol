//ff:func feature=gen-gogin type=generator control=sequence topic=dos-guard
//ff:what blockBodyLimit — middleware.BodyLimit / MultipartLimit / OverrideBodyLimit 등록

package boot

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/middleware"
)

const (
	defaultBodyLimit      = int64(1 << 20) // 1 MiB
	defaultMultipartLimit = int64(32 << 20) // 32 MiB
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

// resolveHTTPLimits computes effective global + per-op limits from
// manifest. Missing manifest → defaults only. Parse failures fall back to
// defaults (SEC validation rule catches bad values at validate-time — the
// generator stays forgiving so a single bad override does not halt codegen).
func resolveHTTPLimits(fs *yongol.Fullstack) (bodyLimit, multipartLimit int64, bodyOverrides, multipartOverrides map[string]int64) {
	bodyLimit = defaultBodyLimit
	multipartLimit = defaultMultipartLimit
	bodyOverrides = map[string]int64{}
	multipartOverrides = map[string]int64{}

	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return
	}
	h := fs.Manifest.Backend.HTTP
	if h.BodyLimit != "" {
		if n, err := middleware.ParseSize(h.BodyLimit); err == nil {
			bodyLimit = n
		}
	}
	if h.MultipartLimit != "" {
		if n, err := middleware.ParseSize(h.MultipartLimit); err == nil {
			multipartLimit = n
		}
	}

	// Build operationId → "METHOD PATH" index from OpenAPI doc.
	opToRoute := buildOperationRouteIndex(fs)
	for opID, ov := range h.Overrides {
		route, ok := opToRoute[opID]
		if !ok {
			continue
		}
		if ov.BodyLimit != "" {
			if n, err := middleware.ParseSize(ov.BodyLimit); err == nil {
				bodyOverrides[route] = n
			}
		}
		if ov.MultipartLimit != "" {
			if n, err := middleware.ParseSize(ov.MultipartLimit); err == nil {
				multipartOverrides[route] = n
			}
		}
	}
	return
}

// buildOperationRouteIndex walks the OpenAPI doc and maps each
// operationId to its gin route key ("METHOD /path"). Returns empty map
// when the doc is nil so callers can still iterate safely.
func buildOperationRouteIndex(fs *yongol.Fullstack) map[string]string {
	idx := map[string]string{}
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return idx
	}
	for path, pi := range fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pi.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			idx[op.OperationID] = method + " " + openAPIPathToGin(path)
		}
	}
	return idx
}

// sortedKeys returns the keys of m sorted deterministically for codegen.
func sortedKeys(m map[string]int64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

