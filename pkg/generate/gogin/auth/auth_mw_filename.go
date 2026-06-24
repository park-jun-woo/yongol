//ff:func feature=gen-gogin type=util control=sequence
//ff:what authMwFileName — strict 미들웨어 출력 파일명 도메인 분리 (단일 사이트는 고정 base.go)

package auth

// authMwFileName derives a strict-middleware output filename (Phase008 §3d).
// Single-site (ident == "") keeps the historical fixed name "<base>.go"; domain
// mode suffixes the domain ident — "<base>_<ident>.go" — so two same-mode
// domains do not overwrite one shared file (1-file-1-func preserved).
func authMwFileName(base, ident string) string {
	if ident == "" {
		return base + ".go"
	}
	return base + "_" + ident + ".go"
}
