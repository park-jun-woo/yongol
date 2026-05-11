//ff:type feature=service type=model
//ff:what Server — StrictServerInterface 구조체 (pgxpool.Pool + sqlc Queries 보관)
package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
)

// Server implements api.StrictServerInterface.
type Server struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
}
