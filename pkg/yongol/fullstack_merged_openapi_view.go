//ff:func feature=orchestrator type=accessor control=sequence
//ff:what MergedOpenAPIView — 모든 도메인 OpenAPI doc 을 합친 단일 doc 으로 채운 fs shallow copy 반환

package yongol

// MergedOpenAPIView returns a shallow copy of fs whose singular OpenAPIDoc is the
// union of every domain's OpenAPI document (Paths + Components), so single-site
// cross-validation rules that consult the FULL operationId / securityScheme set
// against global manifest/hurl config (e.g. openapi_manifest SEC-05 rate_limit
// routability, hurl_openapi route coverage) see all domains at once instead of
// one. operationIds and paths are globally unique under XDO-90, so the path union
// is collision-free; security-scheme/schema names are deduped (last writer wins),
// which is harmless for the membership checks these rules perform.
//
// The shared Ground (already merged in ground.Build) is preserved by the copy.
// OpenAPILines stays nil — there is no single source file for a merged doc, so
// reverse-rule diagnostics fall back to line 0 (cosmetic; only reached on an
// actual violation). In single-site mode there are no domains to merge and the
// caller runs on the real fs instead, so this is only used in domain mode.
func (fs *Fullstack) MergedOpenAPIView() *Fullstack {
	view := *fs
	view.OpenAPIDoc = fs.mergedOpenAPIDoc()
	return &view
}
