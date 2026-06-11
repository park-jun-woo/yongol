//ff:func feature=projectconfig type=util control=sequence
//ff:what RoleFieldOnly — frontend.auth 블록이 role_field 외 키 없이 선언되었는지 판정 (TM-24·XON-60 공유 술어)

package manifest

// RoleFieldOnly reports whether the frontend.auth block declares role_field
// and nothing else. Such a block is the cookie-mode menu role wiring of
// plans/stml/sitemap Phase005: it carries no token contract, so XON-60
// exempts it from the token_field requirement and TM-24 exempts it from the
// cookie-mode conflict WARNING. Both rules consume this single predicate so
// their judgments can never drift. Any token-related key (token_field,
// refresh_field, refresh_op, store) makes the block a token declaration
// again and both rules apply in full.
func (a *FrontendAuth) RoleFieldOnly() bool {
	return a != nil && a.RoleField != "" &&
		a.TokenField == "" && a.RefreshField == "" && a.RefreshOp == "" && a.Store == ""
}
