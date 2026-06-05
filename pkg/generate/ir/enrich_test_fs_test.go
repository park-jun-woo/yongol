//ff:func feature=gen-ir type=test-helper control=sequence
//ff:what enrichTestFS — 테스트 헬퍼

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func enrichTestFS() *yongol.Fullstack {
	return &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name: "courses",
				Columns: map[string]ddl.Column{
					"id":    {Name: "id"},
					"title": {Name: "title"},
				},
				PrimaryKey: []string{"id"},
			},
		},
	}
}
