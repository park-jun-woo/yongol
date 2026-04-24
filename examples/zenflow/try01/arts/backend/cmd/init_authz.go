//ff:func feature=main type=util control=sequence
//ff:what initAuthz — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=d4844539
package main

import (
	"database/sql"
	"github.com/park-jun-woo/ssac/pkg/authz"
	"log/slog"
	"os"
)

func initAuthz(conn *sql.DB) {
	slog.Info("initializing authz")
	opaPath := os.Getenv("OPA_POLICY_PATH")
	if opaPath == "" {
		slog.Error("OPA_POLICY_PATH is required")
		os.Exit(1)
	}
	if _, err := os.Stat(opaPath); err != nil {
		slog.Error("OPA_POLICY_PATH not accessible", "path", opaPath, "err", err)
		os.Exit(1)
	}
	if err := authz.Init(conn, []authz.OwnershipMapping{
			{Resource: "workflow", Table: "workflows", Column: "org_id"},
			{Resource: "execution_log", Table: "execution_logs", Column: "org_id"},
	}); err != nil {
		slog.Error("authz init", "err", err)
		os.Exit(1)
	}
}
