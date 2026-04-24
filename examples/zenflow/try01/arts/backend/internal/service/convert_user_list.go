//ff:func feature=service type=util control=iteration dimension=1 topic=response-serialize
//ff:what convertUserList — []db.User → []api.User 변환
//ff:checked llm=yongol-gen hash=7562cf12
package service

import (
	"github.com/example/zenflow_try01/internal/api"
	"github.com/example/zenflow_try01/internal/db"
)

func convertUserList(rows []db.User) ([]api.User, error) {
	result := make([]api.User, len(rows))
	for i, row := range rows {
		item, err := convertUser(row)
		if err != nil {
			return nil, err
		}
		result[i] = *item
	}
	return result, nil
}
