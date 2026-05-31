//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBindNestJSDispatch — bindNestJS 전 family 디스패치 + unsupported 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindNestJSDispatch_ZeroCov(t *testing.T) {
	fams := []typemap.PGFamily{
		typemap.FamilyInteger, typemap.FamilyFloat, typemap.FamilyString,
		typemap.FamilyBoolean, typemap.FamilyUUID, typemap.FamilyNumeric,
		typemap.FamilyTimestampTZ, typemap.FamilyTimestamp, typemap.FamilyDate,
		typemap.FamilyInet, typemap.FamilyInterval, typemap.FamilyJSONB,
		typemap.FamilyBytea, typemap.FamilyEnum, typemap.FamilyArray,
	}
	for _, f := range fams {
		_ = bindNestJS(f, ir.BindOpts{NotNull: true, ElementHead: "TEXT"})
		_ = bindNestJS(f, ir.BindOpts{NotNull: false, ElementHead: "TEXT"})
	}
	// unsupported default branch.
	if bindNestJS(typemap.FamilyUnsupported, ir.BindOpts{}).Supported {
		t.Errorf("unsupported family should not be Supported")
	}
}
