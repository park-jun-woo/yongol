//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what idxFor — DDL/component/SSaC/Ground 으로 production buildEntityIndex 호출해 entityIndex 생성

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// idxFor builds an entityIndex via the production buildEntityIndex from the
// given DDL tables, component schemas, SSaC funcs and Ground var-type map.
func idxFor(tables []ddl.Table, schemas openapi3.Schemas, funcs []ssac.ServiceFunc, varTypes map[string]string) *entityIndex {
	fs := buildCanonicalFS(tables, schemas, nil, funcs, varTypes)
	return buildEntityIndex(fs)
}
