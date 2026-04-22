//ff:func feature=chain type=util control=iteration dimension=2
//ff:what traceFuncSpecs finds func specs referenced by @call sequences.
package chain

import (
	"log/slog"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func traceFuncSpecs(sf *ssac.ServiceFunc, specs []funcspec.FuncSpec, specsDir string) []Link {
	callPkgFuncs := map[string]string{} // "pkg.Func" -> pkg
	for _, seq := range sf.Sequences {
		if seq.Type != "call" || seq.Model == "" {
			continue
		}
		parts := strings.SplitN(seq.Model, ".", 2)
		if len(parts) == 2 {
			callPkgFuncs[seq.Model] = parts[0]
		}
	}

	if len(callPkgFuncs) == 0 {
		slog.Debug("chain.traceFuncSpecs: no @call sequences in SSaC function", "operationId", sf.Name)
		return nil
	}

	var links []Link
	for callRef, pkg := range callPkgFuncs {
		parts := strings.SplitN(callRef, ".", 2)
		funcName := ""
		if len(parts) == 2 {
			funcName = parts[1]
		}
		if link, ok := findFuncSpecLink(callRef, pkg, funcName, specs, specsDir); ok {
			links = append(links, link)
		}
	}
	return links
}
