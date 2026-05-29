//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what collectSensitiveKeys — DDL `-- @sensitive` 컬럼명을 sorted 리스트로 수집

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectSensitiveKeys(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{Columns: map[string]ddl.Column{
				"ssn":   {Sensitive: true},
				"email": {Sensitive: false},
			}},
			{Columns: map[string]ddl.Column{
				"api_key": {Sensitive: true},
				"ssn":     {Sensitive: true}, // duplicate across tables
			}},
		},
	}
	got := collectSensitiveKeys(fs)
	// sorted + de-duplicated
	if strings.Join(got, ",") != "api_key,ssn" {
		t.Errorf("collectSensitiveKeys = %v, want [api_key ssn]", got)
	}
}

func TestCollectSensitiveKeys_None(t *testing.T) {
	got := collectSensitiveKeys(&yongol.Fullstack{})
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}
