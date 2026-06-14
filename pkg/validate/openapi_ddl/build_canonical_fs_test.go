//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=openapi-ddl
//ff:what buildCanonicalFS — canonicalResponseRepr 테스트용 Fullstack 빌더 (DDL/components/operations/SSaC/Ground)

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildCanonicalFS assembles a Fullstack with the given DDL tables, component
// schemas, operations and SSaC functions, then binds a Ground populated with
// the supplied SSaC.var type map. Each op is mounted on its own path.
func buildCanonicalFS(
	tables []ddl.Table,
	schemas openapi3.Schemas,
	ops []canonOp,
	funcs []ssac.ServiceFunc,
	varTypes map[string]string,
) *yongol.Fullstack {
	paths := openapi3.NewPaths()
	for _, o := range ops {
		op := &openapi3.Operation{OperationID: o.opID, Responses: o.resp}
		item := &openapi3.PathItem{}
		switch o.method {
		case "GET":
			item.Get = op
		case "POST":
			item.Post = op
		case "PUT":
			item.Put = op
		case "DELETE":
			item.Delete = op
		case "PATCH":
			item.Patch = op
		}
		paths.Set("/"+o.opID, item)
	}
	fs := &yongol.Fullstack{
		DDLTables:    tables,
		ServiceFuncs: funcs,
		OpenAPIDoc: &openapi3.T{
			Paths:      paths,
			Components: &openapi3.Components{Schemas: schemas},
		},
		OpenAPILines: &oapiparser.LineIndex{
			Operations:       map[string]int{},
			SchemaProperties: map[string]map[string]int{},
			Schemas:          map[string]int{},
		},
	}
	types := map[string]string{}
	for k, v := range varTypes {
		types[k] = v
	}
	fs.SetGround(&rule.Ground{Types: types, Schemas: map[string][]string{}, Flags: map[string]bool{}})
	return fs
}
