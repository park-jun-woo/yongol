//ff:func feature=validate type=util control=iteration dimension=2 topic=hurl-statemachine
//ff:what buildOpIDLookup — "METHOD /path" → operationId 테이블 (hurl entry 매칭용)

package hurl_statemachine

import (
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var (
	reHurlVarKey    = regexp.MustCompile(`\{\{.+?\}\}`)
	reHurlNumericK  = regexp.MustCompile(`^\d+$`)
	reOpenAPIVarKey = regexp.MustCompile(`^\{.+\}$`)
)

// buildOpIDLookup derives a "METHOD /normalised/path" → operationId
// map from the parsed OpenAPI document. Hurl entries are keyed against
// this map using the same normalisation (`:param` for each variable /
// numeric segment) so the check tolerates concrete test IDs.
func buildOpIDLookup(fs *yongol.Fullstack) map[string]string {
	out := map[string]string{}
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return out
	}
	for path, pi := range fs.OpenAPIDoc.Paths.Map() {
		norm := normPath(path, reOpenAPIVarKey, nil)
		for method, op := range pi.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			out[strings.ToUpper(method)+" "+norm] = op.OperationID
		}
	}
	return out
}
