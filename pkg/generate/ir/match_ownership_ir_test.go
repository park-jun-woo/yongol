//ff:func feature=gen-ir type=test control=sequence
//ff:what TestMatchOwnershipIR -- OwnershipMapping resource 매칭 시 OwnershipInfo 구성(+DDL PK 해석), 미스 시 nil 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestMatchOwnershipIR(t *testing.T) {
	mappings := []rego.OwnershipMapping{
		{Resource: "gig", Table: "gigs", Column: "client_id"},
		{Resource: "note", Table: "notes", Column: "owner_id"},
	}

	// with DDLTables -> ResourcePK resolved
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{Name: "notes", PrimaryKey: []string{"id"}},
		},
	}
	info := matchOwnershipIR(fs, mappings, "note")
	if info == nil {
		t.Fatal("expected OwnershipInfo for note")
	}
	if info.Table != "notes" || info.OwnerColumn != "owner_id" || info.ResourcePK != "id" {
		t.Errorf("info = %+v, want Table=notes OwnerColumn=owner_id ResourcePK=id", info)
	}

	// without DDLTables -> ResourcePK empty
	infoNoDDL := matchOwnershipIR(&yongol.Fullstack{}, mappings, "gig")
	if infoNoDDL == nil {
		t.Fatal("expected OwnershipInfo for gig")
	}
	if infoNoDDL.Table != "gigs" || infoNoDDL.OwnerColumn != "client_id" || infoNoDDL.ResourcePK != "" {
		t.Errorf("info = %+v, want Table=gigs OwnerColumn=client_id ResourcePK=empty", infoNoDDL)
	}

	// no matching resource -> nil
	if got := matchOwnershipIR(fs, mappings, "missing"); got != nil {
		t.Errorf("expected nil for unmatched resource, got %+v", got)
	}
}
