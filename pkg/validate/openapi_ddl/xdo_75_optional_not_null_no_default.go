//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-75 — OpenAPI optional + DDL NOT NULL + DEFAULT 없음 → ERROR

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo75OptionalNotNullNoDefault validates XDO-75: if an OpenAPI request field
// is optional (not required) but the corresponding DDL column has NOT NULL with
// no DEFAULT, INSERT will fail at runtime.
func xdo75OptionalNotNullNoDefault(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic

	for opID, fields := range fs.RequestConstraints {
		diags = append(diags, xdo75CheckOpFields(fs, opID, fields)...)
	}
	return diags
}
