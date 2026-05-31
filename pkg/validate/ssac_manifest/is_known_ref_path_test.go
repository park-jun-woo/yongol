//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"
)

func TestIsKnownRefPath(t *testing.T) {
	if !isKnownRefPath("auth.accessTokenTTL") {
		t.Error("auth.accessTokenTTL should be known")
	}
	if isKnownRefPath("bogus.path") {
		t.Error("bogus.path should not be known")
	}
}
