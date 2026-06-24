//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what buildDomainCORSDispatch — 경로 prefix→도메인 allow_origins 분기 isDomainOriginAllowed 소스 + 필요한 import 산출

package boot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildDomainCORSDispatch emits the isDomainOriginAllowed helper for domain
// mode and the imports it needs. The helper maps a request path to the owning
// domain's resolved allow_origins (Decision L): the most specific route_prefix
// wins, the least specific (shortest) prefix is the default fallback so paths
// outside any /api/<domain> prefix (health probes) still get the broadest
// domain's origins. Each domain's origins are resolved via
// DomainConfig.ResolvedAllowOrigins so a domain inherits backend.cors when it
// omits its own cors block. The returned imports always include "slices"
// (slices.Contains) and add "strings" only when ≥2 domains require
// strings.HasPrefix cases.
func buildDomainCORSDispatch(fs *yongol.Fullstack) (string, []string) {
	backend := fs.Manifest.Backend.CORS
	type dom struct {
		prefix  string
		origins []string
	}
	var doms []dom
	for _, name := range fs.DomainNames() {
		cfg := fs.Manifest.Domains[name]
		doms = append(doms, dom{prefix: cfg.RoutePrefix, origins: cfg.ResolvedAllowOrigins(backend)})
	}
	// Most specific prefix first; stable so equal-length prefixes keep the
	// DomainNames() sort (name asc) for deterministic output.
	sort.SliceStable(doms, func(i, j int) bool { return len(doms[i].prefix) > len(doms[j].prefix) })

	var b strings.Builder
	b.WriteString("func isDomainOriginAllowed(origin, path string) bool {\n")
	imports := []string{`"slices"`}
	if len(doms) <= 1 {
		var origins []string
		if len(doms) == 1 {
			origins = doms[0].origins
		}
		b.WriteString(fmt.Sprintf("\treturn slices.Contains(%s, origin)\n", goStringSlice(origins)))
		b.WriteString("}")
		return b.String(), imports
	}
	imports = append(imports, `"strings"`)
	b.WriteString("\tswitch {\n")
	for _, d := range doms[:len(doms)-1] {
		b.WriteString(fmt.Sprintf("\tcase strings.HasPrefix(path, %q):\n", d.prefix))
		b.WriteString(fmt.Sprintf("\t\treturn slices.Contains(%s, origin)\n", goStringSlice(d.origins)))
	}
	last := doms[len(doms)-1]
	b.WriteString("\tdefault:\n")
	b.WriteString(fmt.Sprintf("\t\treturn slices.Contains(%s, origin)\n", goStringSlice(last.origins)))
	b.WriteString("\t}\n}")
	return b.String(), imports
}
