//ff:func feature=gen-fastapi type=test control=selection
//ff:what TestBindFastAPI — PGFamily 별 디스패치 검증 (모든 family + default unsupported)

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindFastAPI(t *testing.T) {
	cases := []struct {
		family    typemap.PGFamily
		opts      ir.BindOpts
		supported bool
	}{
		{typemap.FamilyInteger, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyFloat, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyString, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyBoolean, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyUUID, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyNumeric, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyTimestampTZ, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyTimestamp, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyDate, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyInet, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyInterval, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyJSONB, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyBytea, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyEnum, ir.BindOpts{NotNull: true}, true},
		{typemap.FamilyArray, ir.BindOpts{NotNull: true, ElementHead: "BIGINT"}, true},
		{typemap.PGFamily(999), ir.BindOpts{}, false},
	}
	for _, c := range cases {
		b := bindFastAPI(c.family, c.opts)
		if b.Supported != c.supported {
			t.Errorf("family %v: Supported = %v, want %v", c.family, b.Supported, c.supported)
		}
		if c.supported && b.Family != c.family {
			t.Errorf("family %v: dispatched Family = %v", c.family, b.Family)
		}
	}
}
