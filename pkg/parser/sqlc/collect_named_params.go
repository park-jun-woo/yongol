//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what collectNamedParams — 한 줄에서 @name 및 sqlc.arg(name) 패턴을 찾아 paramSet 에 누적
package sqlc

import "github.com/park-jun-woo/yongol/pkg/util/caseconv"

// collectNamedParams scans a single SQL source line and records every named
// parameter reference (@name or sqlc.arg(name)) in the caller-supplied set.
func collectNamedParams(line string, paramSet map[string]bool) {
	for _, match := range namedParamRe.FindAllStringSubmatch(line, -1) {
		paramSet[caseconv.SnakeToPascalSqlc(match[1])] = true
	}
	for _, match := range sqlcArgRe.FindAllStringSubmatch(line, -1) {
		paramSet[caseconv.SnakeToPascalSqlc(match[1])] = true
	}
}
