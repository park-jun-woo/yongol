//ff:func feature=validate type=rule control=sequence topic=ssac-manifest
//ff:what Run — SSaC↔Manifest 교차 검증 실행 (XNS-*)
package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔Manifest cross-validation rules.
//
// XSA-70/71/72 are ERROR rules: SSaC uses session/cache/file built-ins but
// the manifest omits the matching backend. Without them, validate passes
// and generate panics on a nil dereference in block_*_init. XNS-56
// already covers the queue case with the same semantics, so no separate
// XSA-73 is emitted.
//
// XSA-74..77 are WARNING rules: the manifest declares a backend that no
// SSaC function uses — hygiene flag to avoid wiring unused infrastructure.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xns48CurrentUserClaims(fs)...)
	diags = append(diags, xns49CurrentUserField(fs)...)
	diags = append(diags, xns56QueueRequired(fs)...)
	diags = append(diags, xns57MemoryTxPublish(fs)...)
	diags = append(diags, xns73JwtCallClaims(fs)...)
	diags = append(diags, xsa70SessionBackendRequired(fs)...)
	diags = append(diags, xsa71CacheBackendRequired(fs)...)
	diags = append(diags, xsa72FileBackendRequired(fs)...)
	diags = append(diags, xsa74SessionBackendUnused(fs)...)
	diags = append(diags, xsa75CacheBackendUnused(fs)...)
	diags = append(diags, xsa76FileBackendUnused(fs)...)
	diags = append(diags, xsa77QueueBackendUnused(fs)...)
	return diags
}
