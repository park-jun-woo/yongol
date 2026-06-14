//ff:type feature=validate type=model topic=openapi-ddl
//ff:what entityIndex — canonical 응답 검증용 DDL/component/SSaC 인덱스 묶음

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// entityIndex bundles the lookups XDO-11/12 need to resolve an OpenAPI response
// to a canonical entity (DDL table ↔ component) and to read SSaC var types.
type entityIndex struct {
	g              *rule.Ground
	tables         []ddl.Table
	tableExists    map[string]bool
	schemaForTable map[string]string // DDL table name → component schema name
	funcByName     map[string]*ssac.ServiceFunc
}
