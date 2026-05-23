//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XFS-39 — @call references an existing func implementation

package ssac_func

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs39CallToFuncSpec validates XFS-39: every @call must reference an
// existing func implementation.
func xfs39CallToFuncSpec(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	specs := g.Lookup["Func.spec"]
	if specs == nil {
		specs = map[string]bool{}
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" {
				continue
			}
			if !strings.Contains(seq.Model, ".") {
				continue
			}
			key := normalizedCallKey(seq.Model)
			if !specs[key] {
				diags = append(diags, diagnostic.Diagnostic{
					File:        fn.FileName,
					Line:        seq.Line,
					Phase:       diagnostic.PhaseValidate,
					Level:       diagnostic.LevelError,
					Message:     "[XFS-39] @call references non-existent func spec " + seq.Model,
					Advice:      xfs39Advice(seq.Model, fs),
					OperationID: fn.Name,
				})
			}
		}
	}
	return diags
}

// xfs39Advice returns contextual advice for a missing @call target. When the
// call targets a builtin ssac package, it lists the available functions in
// that package (merged from YongolPkgSpecs and Ground). Otherwise it falls
// back to the generic "define under pkg/" message.
func xfs39Advice(model string, fs *yongol.Fullstack) string {
	idx := strings.IndexByte(model, '.')
	if idx <= 0 {
		return "Define function " + model + " under pkg/"
	}
	pkg := model[:idx]
	if !builtinPackages[pkg] {
		return "Define function " + model + " under pkg/"
	}
	names := collectBuiltinFuncNames(pkg, fs)
	if len(names) == 0 {
		return "Package " + pkg + " is a builtin ssac package but no functions were loaded. Check ssac/pkg installation."
	}
	return "Available " + pkg + " functions: " + strings.Join(names, ", ")
}

// collectBuiltinFuncNames merges function names from YongolPkgSpecs and
// Ground Func.spec for a given builtin package. Returns sorted, deduplicated
// PascalCase names.
func collectBuiltinFuncNames(pkg string, fs *yongol.Fullstack) []string {
	seen := map[string]bool{}
	for _, name := range builtinFuncNames(pkg, fs.YongolPkgSpecs) {
		seen[name] = true
	}
	// Also include functions registered in Ground (e.g. auth.issueToken added
	// by populateFunc when manifest has auth.claims).
	if g := fs.Ground(); g != nil {
		if funcSpecs := g.Lookup["Func.spec"]; funcSpecs != nil {
			prefix := pkg + "."
			for key := range funcSpecs {
				if strings.HasPrefix(key, prefix) {
					camel := key[len(prefix):]
					seen[ucFirst(camel)] = true
				}
			}
		}
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
