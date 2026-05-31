//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestScanValidateLineForTerminator(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		inSingle     bool
		wantDone     bool
		wantInSingle bool
	}{
		{"unquoted semicolon", "VALUES (1);", false, true, false},
		{"no semicolon", "VALUES (1)", false, false, false},
		{"semicolon inside quote ignored", "'a;b'", false, false, false},
		{"open quote carries over", "'unterminated", false, false, true},
		{"close quote from carried state", "more'", true, false, false},
		{"escaped doubled quote stays in literal", "'it''s';", false, true, false},
		{"semicolon while inside single is ignored", ";", true, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done, inSingle := scanValidateLineForTerminator(tt.line, tt.inSingle)
			if done != tt.wantDone || inSingle != tt.wantInSingle {
				t.Errorf("scanValidateLineForTerminator(%q,%v) = (%v,%v), want (%v,%v)",
					tt.line, tt.inSingle, done, inSingle, tt.wantDone, tt.wantInSingle)
			}
		})
	}
}
