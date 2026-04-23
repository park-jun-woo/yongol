//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what fillMissingFields — fills empty RequestFields/ResponseFields from package-level types and propagates collection errors as Diagnostics
package funcspec

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// fillMissingFields fills empty RequestFields/ResponseFields from
// companion struct files in the same directory.
//
// Diagnostics returned by collectPackageTypes are deduplicated per directory
// via seenDir to avoid appending them more than once for the same directory.
func fillMissingFields(specs []FuncSpec, specDirs []string) []diagnostic.Diagnostic {
	cache := make(map[string]map[string][]Field)
	seenDir := make(map[string]struct{})
	var diags []diagnostic.Diagnostic
	for i := range specs {
		if len(specs[i].RequestFields) > 0 && len(specs[i].ResponseFields) > 0 {
			continue
		}
		dir := specDirs[i]
		var typeMap map[string][]Field
		typeMap, diags = loadTypeMapForDir(dir, cache, seenDir, diags)
		fillSpecFromTypeMap(&specs[i], typeMap)
	}
	return diags
}
