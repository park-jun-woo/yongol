//ff:type feature=gen-react type=model
//ff:what refreshPlan — bearer 401→refresh→재시도 흐름에 필요한 refresh op 결정값

package react

// refreshPlan carries everything writeRefreshFlow needs to emit the
// single-flight 401→refresh→retry flow: the refresh operation identity
// (declared via frontend.auth.refresh_op or structurally inferred), the
// token fields committed back into the session store, and how the stored
// refresh token travels in the request.
type refreshPlan struct {
	opID         string // refresh operationId (declared or inferred)
	method       string // upper-case HTTP method, e.g. "POST"
	path         string // OpenAPI path key, e.g. "/auth/refresh"
	tokenField   string // frontend.auth.token_field — 2xx property committed as auth.token
	refreshField string // frontend.auth.refresh_field — 2xx property committed as auth.refresh
	// bodyKey is the requestBody property that carries the stored refresh
	// token (matched by refreshField name). Empty means the op takes no
	// such body property (e.g. cookie-carried refresh) and the call sends
	// no body.
	bodyKey string
}
