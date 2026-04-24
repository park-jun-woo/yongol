//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what authzHelperInitAuthzFactory — OPA authz.Init 호출 헬퍼 (OwnershipMapping 임베드) 소스 생성

package boot

import "strings"

// authzHelperInitAuthzFactory returns the top-level initAuthz(policyPath)
// helper source with the OwnershipMapping literal embedded. Extracted from
// main() so the function stays under Q3's 100-line sequence limit. Exits with
// os.Exit(1) on any precondition failure matching the previous inline
// behavior.
//
// Phase002 (ssac/purify) — ssac/pkg/authz is DB-free so the helper no longer
// takes *sql.DB. The policy path is the single external dependency; ownership
// lookups happen in handler code via user sqlc queries (Phase003).
func authzHelperInitAuthzFactory(mappings []string) string {
	var sb strings.Builder
	sb.WriteString("func initAuthz(policyPath string) {\n")
	sb.WriteString("\tslog.Info(\"initializing authz\")\n")
	sb.WriteString("\tif policyPath == \"\" {\n")
	sb.WriteString("\t\tslog.Error(\"OPA_POLICY_PATH is required\")\n")
	sb.WriteString("\t\tos.Exit(1)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tif _, err := os.Stat(policyPath); err != nil {\n")
	sb.WriteString("\t\tslog.Error(\"OPA_POLICY_PATH not accessible\", \"path\", policyPath, \"err\", err)\n")
	sb.WriteString("\t\tos.Exit(1)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tif err := authz.Init(policyPath, []authz.OwnershipMapping{\n")
	for _, m := range mappings {
		sb.WriteString("\t" + m + "\n")
	}
	sb.WriteString("\t}); err != nil {\n")
	sb.WriteString("\t\tslog.Error(\"authz init\", \"err\", err)\n")
	sb.WriteString("\t\tos.Exit(1)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}")
	return sb.String()
}
