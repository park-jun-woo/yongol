//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithFuncs(funcs ...ssac.ServiceFunc) *yongol.Fullstack {
	return &yongol.Fullstack{ServiceFuncs: funcs}
}

func callSeq(model string) ssac.Sequence {
	return ssac.Sequence{Type: "call", Model: model}
}

func TestNormalizeCallKey(t *testing.T) {
	tests := []struct{ in, want string }{
		{"session.GetUser", "session.getUser"},
		{"cache.Set", "cache.set"},
		{"VerifyPassword", "verifyPassword"},
	}
	for _, tt := range tests {
		if got := normalizeCallKey(tt.in); got != tt.want {
			t.Errorf("normalizeCallKey(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestUsesCache(t *testing.T) {
	yes := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	if !usesCache(yes) {
		t.Error("expected cache use detected")
	}
	no := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("session.Get")}})
	if usesCache(no) {
		t.Error("session.Get should not count as cache use")
	}
	if usesCache(fsWithFuncs()) {
		t.Error("empty fs should not use cache")
	}
}

func TestUsesSession(t *testing.T) {
	yes := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("session.Put")}})
	if !usesSession(yes) {
		t.Error("expected session use detected")
	}
	no := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	if usesSession(no) {
		t.Error("cache.Get should not count as session use")
	}
}

func TestUsesFile(t *testing.T) {
	if !usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})) {
		t.Error("file. prefix expected")
	}
	if !usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("storage.Upload")}})) {
		t.Error("storage. prefix expected")
	}
	if usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})) {
		t.Error("cache.Get should not count as file use")
	}
	// non-call sequence ignored.
	if usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get", Model: "file.X"}}})) {
		t.Error("non-call seq should be ignored")
	}
}

func TestUsesQueue(t *testing.T) {
	// publish sequence.
	if !usesQueue(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}})) {
		t.Error("publish should count")
	}
	// subscribe func.
	if !usesQueue(fsWithFuncs(ssac.ServiceFunc{Subscribe: &ssac.SubscribeInfo{}})) {
		t.Error("subscribe should count")
	}
	if usesQueue(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})) {
		t.Error("no queue usage expected")
	}
}

func TestUsesCurrentUser(t *testing.T) {
	// auth sequence.
	if !usesCurrentUser([]ssac.ServiceFunc{{Sequences: []ssac.Sequence{{Type: "auth"}}}}) {
		t.Error("auth seq should count")
	}
	// currentUser. input.
	withInput := []ssac.ServiceFunc{{Sequences: []ssac.Sequence{
		{Type: "post", Inputs: map[string]string{"owner": "currentUser.ID"}},
	}}}
	if !usesCurrentUser(withInput) {
		t.Error("currentUser. input should count")
	}
	// no usage.
	if usesCurrentUser([]ssac.ServiceFunc{{Sequences: []ssac.Sequence{
		{Type: "post", Inputs: map[string]string{"x": "body.X"}},
	}}}) {
		t.Error("plain input should not count")
	}
}

func TestHasTxBoundPublish(t *testing.T) {
	// mutation + publish → true.
	fn := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "post"}, {Type: "publish"}}}
	if !hasTxBoundPublish(fn) {
		t.Error("post + publish should be tx-bound")
	}
	// publish without mutation → false.
	if hasTxBoundPublish(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get"}, {Type: "publish"}}}) {
		t.Error("publish w/o mutation should be false")
	}
	// mutation without publish → false.
	if hasTxBoundPublish(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "put"}}}) {
		t.Error("mutation w/o publish should be false")
	}
}

func TestIsKnownRefPath(t *testing.T) {
	if !isKnownRefPath("auth.accessTokenTTL") {
		t.Error("auth.accessTokenTTL should be known")
	}
	if isKnownRefPath("bogus.path") {
		t.Error("bogus.path should not be known")
	}
}

func TestQueueBackend(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Queue: &manifest.QueueBackend{Backend: "postgres"}}}
	if got := queueBackend(fs); got != "postgres" {
		t.Errorf("got %q, want postgres", got)
	}
	// nil manifest → "".
	if got := queueBackend(&yongol.Fullstack{}); got != "" {
		t.Errorf("nil manifest: %q", got)
	}
	// nil Queue → "".
	if got := queueBackend(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}); got != "" {
		t.Errorf("nil queue: %q", got)
	}
}

func TestSortedClaimFields(t *testing.T) {
	got := sortedClaimFields(map[string]bool{"role": true, "id": true, "email": true})
	if got != "email, id, role" {
		t.Errorf("got %q, want 'email, id, role'", got)
	}
	if got := sortedClaimFields(map[string]bool{}); got != "" {
		t.Errorf("empty → %q", got)
	}
}
