//ff:func feature=main type=util control=sequence
//ff:what initAuthz — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=d63fd15f
package main

import (
	"github.com/park-jun-woo/ssac/pkg/authz"
	"log/slog"
	"os"
)

func initAuthz(policyPath string) {
	slog.Info("initializing authz")
	if policyPath == "" {
		slog.Error("OPA_POLICY_PATH is required")
		os.Exit(1)
	}
	if _, err := os.Stat(policyPath); err != nil {
		slog.Error("OPA_POLICY_PATH not accessible", "path", policyPath, "err", err)
		os.Exit(1)
	}
	if err := authz.Init(policyPath, []authz.OwnershipMapping{
	}); err != nil {
		slog.Error("authz init", "err", err)
		os.Exit(1)
	}
}
