//ff:type feature=service type=model
//ff:what Server — StrictServerInterface 구조체 (pgxpool.Pool/Queries 보관, auth 활성 시 RefreshStore 추가)
package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/park-jun-woo/zenflow/internal/db"
	"github.com/park-jun-woo/ssac/pkg/auth"
)

// Server implements api.StrictServerInterface.
type Server struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
	RefreshStore *auth.RefreshStore
}
