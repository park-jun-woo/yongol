//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertTemplateList — []db.Template → []api.Template 변환
//ff:checked llm=yongol-gen hash=12901403
package service

import (
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

func convertTemplateList(rows []db.Template) ([]api.Template, error) {
	result := make([]api.Template, len(rows))
	for i, row := range rows {
		item, err := convertTemplate(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
