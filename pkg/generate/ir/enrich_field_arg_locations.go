//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what enrichFieldArgLocations -- OpenAPI path/query 분류 기반으로 모든 Op 의 FieldArg.Location 세팅

package ir

// enrichFieldArgLocations sets FieldArg.Location for every op's arguments
// based on the OpenAPI path/query classification. Arguments with Source
// "request" are classified as path/query/body; Source "currentUser" maps to
// LocUser; quoted literals to LocLiteral; everything else to LocVar.
func enrichFieldArgLocations(ops []Op, pathParams, queryParams map[string]bool) {
	for i := range ops {
		for _, args := range collectFieldArgSlices(&ops[i]) {
			for j := range *args {
				fa := &(*args)[j]
				classifyFieldArgLocation(fa, pathParams, queryParams)
			}
		}
	}
}
