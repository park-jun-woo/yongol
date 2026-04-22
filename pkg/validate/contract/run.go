//ff:func feature=validate-contract type=rule control=sequence
//ff:what Run — arts 디렉토리 preserved 파일 대상 PRV-01/02 + PRV-10~17 오케스트레이션

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/contract"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes every preserve-scoped rule against every preserved
// `.go` file found under artsDir:
//
//   - PRV-01  signature drift
//   - PRV-02  external symbol drift
//   - PRV-10  panic() reintroduced
//   - PRV-11  currentUser unsafe type assertion
//   - PRV-12  json/yaml Unmarshal error ignored
//   - PRV-13  sql.Scan error ignored
//   - PRV-14  slice[0] without len guard
//   - PRV-15  map/index access dereferenced without comma-ok
//   - PRV-16  Get*/Find* return value dereferenced without nil guard
//   - PRV-17  resource acquired without matching defer Close
//
// Preconditions:
//   - fs must have been built by yongol.ParseAll and had its Ground
//     populated (pkg/validate.Validate handles this).
//   - artsDir is the root of the generated artifact tree (the second
//     argument to `yongol generate`).
//
// When artsDir does not exist or contains no preserved files the
// function returns an empty diagnostic slice — the caller is
// responsible for skipping contract validation entirely when no arts
// directory was supplied (see cmd/yongol/validate_cmd.go).
func Run(fs *yongol.Fullstack, artsDir string) []diagnostic.Diagnostic {
	if fs == nil || artsDir == "" {
		return nil
	}
	paths, err := contract.CollectPreserved(artsDir)
	if err != nil {
		return nil
	}
	if len(paths) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	diags = append(diags, prv01SignatureDrift(fs, paths)...)
	diags = append(diags, prv02ExternalSymbolDrift(fs, paths)...)
	diags = append(diags, prv10PreservedPanic(paths)...)
	diags = append(diags, prv11PreservedCurrentUserAssertion(paths)...)
	diags = append(diags, prv12PreservedUnmarshalErr(paths)...)
	diags = append(diags, prv13PreservedScanErr(paths)...)
	diags = append(diags, prv14PreservedSliceBounds(paths)...)
	diags = append(diags, prv15PreservedMapAccess(paths)...)
	diags = append(diags, prv16PreservedNilDeref(paths)...)
	diags = append(diags, prv17PreservedMissingClose(paths)...)
	return diags
}
