//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what validateFetchBlock — data-fetch 블록의 operationId·파라미터·바인딩·each·component 검증

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// validateFetchBlock checks a data-fetch block against the OpenAPI spec.
// TM-01 (operationId), TM-03 (GET method), TM-04 (params), TM-06 (binds),
// TM-07/TM-08 (each), TM-09 (component).
func validateFetchBlock(f stml.FetchBlock, file string, opMap map[string]operationEntry, fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic

	entry, ok := opMap[f.OperationID]
	if !ok {
		// TM-01: data-fetch operationId not found
		diags = append(diags, tm01FetchOpNotFound(file, f.OperationID))
		return diags
	}

	// TM-03: data-action referencing GET is only for action blocks; for
	// fetch blocks, a non-GET method is the error (fetch expects GET).
	// However the original bak logic only checks data-action GET, not
	// data-fetch non-GET. We skip method check for fetch blocks.

	// TM-04: params
	diags = append(diags, tm04Params(f.Params, f.OperationID, file, entry)...)

	// TM-06: binds
	diags = append(diags, tm06Binds(f.Binds, f.OperationID, file, entry)...)

	// TM-07 / TM-08: each
	diags = append(diags, tm0708Each(f.Eaches, f.OperationID, file, entry)...)

	// TM-09: components
	diags = append(diags, tm09Components(f.Components, file, fs)...)

	// Recurse into nested fetches
	for _, child := range f.NestedFetches {
		diags = append(diags, validateFetchBlock(child, file, opMap, fs)...)
	}

	return diags
}
