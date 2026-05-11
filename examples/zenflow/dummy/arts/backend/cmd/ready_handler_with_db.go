//ff:func feature=main type=util control=sequence
//ff:what readyHandlerWithDB — 환경변수 파싱 헬퍼 (실패 시 default 반환)
//ff:checked llm=yongol-gen hash=df87147b
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

func readyHandlerWithDB(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		pingCtx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := pool.Ping(pingCtx); err != nil {
			slog.Warn("readiness: db ping failed", "err", err)
			c.JSON(503, gin.H{"status": "unavailable", "checks": gin.H{"db": "down"}})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "checks": gin.H{"db": "ok"}})
	}
}
