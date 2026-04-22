//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what authzHelperInitAuthzFactory — OPA authz.Init 호출 헬퍼 (OwnershipMapping 임베드) 소스 생성

package boot

import "strings"

// authzHelperInitAuthzFactory returns the top-level initAuthz(conn *sql.DB) helper
// source with the OwnershipMapping literal embedded. Extracted from main() so the
// function stays under Q3's 100-line sequence limit. Exits with os.Exit(1) on any
// precondition failure matching the previous inline behavior.
func authzHelperInitAuthzFactory(mappings []string) string {
	var sb strings.Builder
	sb.WriteString("func initAuthz(conn *sql.DB) {\n")
	sb.WriteString("\tslog.Info(\"initializing authz\")\n")
	sb.WriteString("\topaPath := os.Getenv(\"OPA_POLICY_PATH\")\n")
	sb.WriteString("\tif opaPath == \"\" {\n")
	sb.WriteString("\t\tslog.Error(\"OPA_POLICY_PATH is required\")\n")
	sb.WriteString("\t\tos.Exit(1)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tif _, err := os.Stat(opaPath); err != nil {\n")
	sb.WriteString("\t\tslog.Error(\"OPA_POLICY_PATH not accessible\", \"path\", opaPath, \"err\", err)\n")
	sb.WriteString("\t\tos.Exit(1)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tif err := authz.Init(conn, []authz.OwnershipMapping{\n")
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
