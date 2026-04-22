//ff:func feature=gen-gogin type=generator control=sequence
//ff:what readyHandlerLines — /ready 핸들러 코드 라인 (DDL 유무에 따라 DB ping 또는 정적 200)

package boot

// readyHandlerLines emits the lines that register the /ready endpoint.
// When withDB is true, the handler delegates to the readyHandlerWithDB helper
// (emitted as a sibling top-level func) so main() stays within Q3's 100-line
// sequence limit. Without DB, registers a static 200 matching /health.
func readyHandlerLines(withDB bool) []string {
	if !withDB {
		return []string{
			`r.GET("/ready", func(c *gin.Context) {`,
			`	c.JSON(200, gin.H{"status": "ok"})`,
			`})`,
		}
	}
	return []string{
		`r.GET("/ready", readyHandlerWithDB(conn))`,
	}
}
