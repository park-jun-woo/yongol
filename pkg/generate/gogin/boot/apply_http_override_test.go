//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what applyHTTPOverride — opID/route 가 매칭될 때 body/multipart override 두 맵에 기록

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestApplyHTTPOverride(t *testing.T) {
	opToRoute := map[string]string{"Upload": "POST /upload"}

	t.Run("unknown opID writes nothing", func(t *testing.T) {
		bo := map[string]int64{}
		mo := map[string]int64{}
		applyHTTPOverride(opToRoute, "Missing", pmanifest.HTTPOverride{BodyLimit: "1MiB"}, bo, mo)
		if len(bo) != 0 || len(mo) != 0 {
			t.Errorf("unknown opID should write nothing, got bo=%v mo=%v", bo, mo)
		}
	})

	t.Run("matched opID writes both maps", func(t *testing.T) {
		bo := map[string]int64{}
		mo := map[string]int64{}
		applyHTTPOverride(opToRoute, "Upload", pmanifest.HTTPOverride{BodyLimit: "10MiB", MultipartLimit: "100MiB"}, bo, mo)
		if bo["POST /upload"] != int64(10<<20) {
			t.Errorf("body override = %d, want %d", bo["POST /upload"], int64(10<<20))
		}
		if mo["POST /upload"] != int64(100<<20) {
			t.Errorf("multipart override = %d, want %d", mo["POST /upload"], int64(100<<20))
		}
	})

	t.Run("unparseable size is skipped", func(t *testing.T) {
		bo := map[string]int64{}
		mo := map[string]int64{}
		applyHTTPOverride(opToRoute, "Upload", pmanifest.HTTPOverride{BodyLimit: "garbage"}, bo, mo)
		if len(bo) != 0 {
			t.Errorf("unparseable body limit should be skipped, got %v", bo)
		}
	})
}
