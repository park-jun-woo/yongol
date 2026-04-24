//ff:type feature=service type=model
//ff:what Server — StrictServerInterface 구조체 (DB/Queries 보관, auth 활성 시 RefreshStore 추가)
package service

import (
	"database/sql"
	"github.com/park-jun-woo/zenflow/internal/db"
	"github.com/park-jun-woo/ssac/pkg/auth"
)

// Server implements api.StrictServerInterface.
type Server struct {
	DB      *sql.DB
	Queries *db.Queries
	RefreshStore *auth.RefreshStore
}
