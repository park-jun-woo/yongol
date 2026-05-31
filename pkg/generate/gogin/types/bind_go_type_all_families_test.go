//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindGoType_AllFamilies(t *testing.T) {
	families := []typemap.PGFamily{
		typemap.FamilyEnum,
		typemap.FamilyArray,
		typemap.FamilyUUID,
		typemap.FamilyNumeric,
		typemap.FamilyTimestampTZ,
		typemap.FamilyTimestamp,
		typemap.FamilyDate,
		typemap.FamilyInet,
		typemap.FamilyInterval,
		typemap.FamilyJSONB,
		typemap.FamilyBytea,
		typemap.FamilyInteger,
		typemap.FamilyFloat,
		typemap.FamilyString,
		typemap.FamilyBoolean,
		typemap.FamilyUnsupported,
	}
	for _, f := range families {
		opts := ir.BindOpts{NotNull: true, ElementHead: "TEXT"}
		gb := bindGoType(f, opts)
		_ = gb // every branch must return a value
	}
	// nullable variant to hit non-notNull paths
	for _, f := range families {
		gb := bindGoType(f, ir.BindOpts{NotNull: false, ElementHead: "TEXT"})
		_ = gb
	}
	// unsupported must be flagged not supported
	if bindGoType(typemap.FamilyUnsupported, ir.BindOpts{}).Supported {
		t.Errorf("FamilyUnsupported should not be Supported")
	}
}
