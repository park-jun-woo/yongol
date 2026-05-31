//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCanonicalDDLTableSet(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.table": {"Workflows": true, "AuditLogs": true},
	}})
	set := canonicalDDLTableSet(fs)
	// canonicalTableKey lowercases + snake + singular.
	if !set["workflow"] {
		t.Errorf("expected canonical 'workflow' key, got %v", set)
	}
	if !set["audit_log"] {
		t.Errorf("expected canonical 'audit_log' key, got %v", set)
	}
}
