//ff:func feature=gen-react type=util control=sequence
//ff:what resolveLayoutAuthMode — 레이아웃 data-logout 방출 모드 결정 ("bearer"/"cookie"/""=auth 없음)

package react

// resolveLayoutAuthMode collapses the frontend auth gates into the layout
// emitter's logout wiring mode (page-flow Phase010): "bearer" clears the
// session store, any other present mode (cookie/hybrid) relies on the
// server op, and "" (no auth) skips the emission entirely (TM-38 already
// warned about the dead declaration).
func resolveLayoutAuthMode(hasAuth, bearerAuth bool) string {
	if !hasAuth {
		return ""
	}
	if bearerAuth {
		return "bearer"
	}
	return "cookie"
}
