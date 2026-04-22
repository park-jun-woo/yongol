//ff:func feature=gen-gogin type=generator control=sequence
//ff:what healthHelperReadyWithDB — /ready DB ping 핸들러 생성 헬퍼 소스 반환

package boot

// healthHelperReadyWithDB returns the top-level readyHandlerWithDB(conn *sql.DB)
// gin.HandlerFunc factory source. Performs a DB PingContext with a 2s timeout
// and returns 503 if the ping fails. Extracted from main() so the /ready
// registration is one line and main stays under Q3's 100-line limit.
func healthHelperReadyWithDB() string {
	return `func readyHandlerWithDB(conn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pingCtx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := conn.PingContext(pingCtx); err != nil {
			slog.Warn("readiness: db ping failed", "err", err)
			c.JSON(503, gin.H{"status": "unavailable", "checks": gin.H{"db": "down"}})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "checks": gin.H{"db": "ok"}})
	}
}`
}
