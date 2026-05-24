//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what buildSensitiveColumnIndex — empty/no-sensitive/single/lowercase 케이스 검증

package query

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestBuildSensitiveColumnIndex(t *testing.T) {
	t.Run("empty tables returns empty map", func(t *testing.T) {
		got := buildSensitiveColumnIndex(nil)
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %+v", got)
		}
	})

	t.Run("table without sensitive columns skipped", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "Users",
				Columns: map[string]ddl.Column{
					"id":   {Name: "id", RawType: "BIGINT", Sensitive: false},
					"name": {Name: "name", RawType: "TEXT", Sensitive: false},
				},
			},
		}
		got := buildSensitiveColumnIndex(tables)
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %+v", got)
		}
	})

	t.Run("single table with sensitive columns", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "Users",
				Columns: map[string]ddl.Column{
					"id":       {Name: "id", RawType: "BIGINT", Sensitive: false},
					"email":    {Name: "email", RawType: "TEXT", Sensitive: true},
					"password": {Name: "password", RawType: "TEXT", Sensitive: true},
				},
			},
		}
		got := buildSensitiveColumnIndex(tables)
		want := map[string][]string{
			"users": {"email", "password"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %+v, want %+v", got, want)
		}
	})

	t.Run("table name lowercased", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "UserProfile",
				Columns: map[string]ddl.Column{
					"ssn": {Name: "ssn", RawType: "TEXT", Sensitive: true},
				},
			},
		}
		got := buildSensitiveColumnIndex(tables)
		if _, ok := got["userprofile"]; !ok {
			t.Fatalf("expected lowercase key 'userprofile', got keys %+v", got)
		}
	})
}
