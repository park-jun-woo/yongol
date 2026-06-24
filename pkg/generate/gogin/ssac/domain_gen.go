//ff:type feature=gen-gogin type=model
//ff:what domainGen — internal/service 코드젠의 도메인별 emission 컨텍스트 (api alias suffix + 함수명 prefix)

package ssac

// domainGen carries per-domain emission context for the shared internal/service
// package. The zero value (both fields empty) is single-site mode: the api
// import stays the bare "<module>/internal/api" with no alias and converter
// names keep their plain convert<Name> form. In domain mode ApiSuffix is
// sanitizeDomainName(name) — the alias target becomes
// `api "<module>/internal/api_<ApiSuffix>"`, resolving every body `api.X` with
// zero edits — and FuncPrefix is domainTitle(name), making converters
// convert<FuncPrefix><Name> so two domains that share a schema name do not
// redeclare the same function. Domain mode is detected by ApiSuffix != "".
type domainGen struct {
	ApiSuffix  string
	FuncPrefix string
}
