//ff:func feature=orchestrator type=loader control=sequence dimension=1
//ff:what 도메인별 단일 파일 SSOT 경로 존재 여부를 SSOTPresence 로 판정
package yongol

import "os"

// probePresence reports whether a single-file SSOT exists at path. It mirrors
// DetectSSOTs' single-file branch (os.Stat existence check): present ⇒
// SSOTPopulated, missing ⇒ SSOTAbsent. Used for per-domain OpenAPI paths.
func probePresence(path string) SSOTPresence {
	if _, err := os.Stat(path); err == nil {
		return SSOTPopulated
	}
	return SSOTAbsent
}
