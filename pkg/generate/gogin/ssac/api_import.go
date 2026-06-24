//ff:func feature=gen-gogin type=util control=sequence
//ff:what apiImportLine — internal/service 파일의 api 패키지 import 라인 (도메인 모드는 api alias)

package ssac

// apiImportLine returns the import spec an internal/service file uses to bring
// in the oapi-codegen api package. Single-site (apiSuffix == "") returns the
// plain `"<module>/internal/api"`. Domain mode aliases the per-domain package
// literally to `api` — `api "<module>/internal/api_<apiSuffix>"` — so every bare
// `api.X` reference in the shared service files resolves to the owning domain's
// types without any body edits. The caller is responsible for surrounding tab /
// newline within an import block.
func apiImportLine(modulePath, apiSuffix string) string {
	if apiSuffix == "" {
		return `"` + modulePath + `/internal/api"`
	}
	return `api "` + modulePath + `/internal/api_` + apiSuffix + `"`
}
