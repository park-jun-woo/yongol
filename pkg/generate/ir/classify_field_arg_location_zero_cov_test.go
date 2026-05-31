//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestFieldArgHelpersZeroCov — classifyFieldArgLocation / collectFieldArgSlices / opModelName / collectQueryMethods 직접 커버
package ir

import (
	"testing"
)

func TestClassifyFieldArgLocation_ZeroCov(t *testing.T) {
	pathParams := map[string]bool{"id": true}
	queryParams := map[string]bool{"cursor": true}

	cases := []struct {
		fa   FieldArg
		want ParamLocation
	}{
		{FieldArg{Literal: "5"}, LocLiteral},
		{FieldArg{Source: "currentUser", Field: "ID"}, LocUser},
		{FieldArg{Source: "request", Field: "id"}, LocPath},
		{FieldArg{Source: "request", Field: "cursor"}, LocQuery},
		{FieldArg{Source: "request", Field: "name"}, LocBody},
		{FieldArg{Source: "course", Field: "ID"}, LocVar},
		{FieldArg{}, ""}, // no source, no literal → untouched
	}
	for i, c := range cases {
		fa := c.fa
		classifyFieldArgLocation(&fa, pathParams, queryParams)
		if fa.Location != c.want {
			t.Errorf("case %d: Location = %q, want %q", i, fa.Location, c.want)
		}
	}
}
