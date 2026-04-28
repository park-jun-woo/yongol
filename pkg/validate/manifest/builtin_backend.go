//ff:type feature=validate type=model topic=manifest-infra
//ff:what builtinBackend — XNC/XNS/XNQ/XNA-90 공용 backend 구성 carrier

package manifest

// builtinBackend is a small carrier used by validateBuiltinBackend so the
// signature stays stable for the 4 symmetrical rules (cache / session /
// queue / auth.refresh).
type builtinBackend struct {
	Present bool
	Backend string
}
