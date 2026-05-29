//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what jsonbConvertRHS 단위 테스트 (required → 값, nullable → &local, no match → false)

package ssac

import "testing"

func TestJsonbConvertRHS(t *testing.T) {
	aliases := []jsonbFieldAlias{
		{jsonName: "meta", apiField: "Meta", dbField: "Meta", localVar: "metaMap"},
	}
	cases := []struct {
		name     string
		jsonName string
		required bool
		wantRHS  string
		wantOK   bool
	}{
		{"required match", "meta", true, "metaMap", true},
		{"nullable match wraps", "meta", false, "&metaMap", true},
		{"no match", "other", true, "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotRHS, gotOK := jsonbConvertRHS(tc.jsonName, tc.required, aliases)
			if gotRHS != tc.wantRHS || gotOK != tc.wantOK {
				t.Errorf("jsonbConvertRHS(%q,%v) = (%q,%v), want (%q,%v)", tc.jsonName, tc.required, gotRHS, gotOK, tc.wantRHS, tc.wantOK)
			}
		})
	}
}
