//ff:func feature=gen-react type=helper control=sequence dimension=1
//ff:what resolveManifestRefreshField -- manifest에서 frontend.auth.refresh_field를 안전하게 추출한다

package react

import "github.com/park-jun-woo/yongol/pkg/yongol"

func resolveManifestRefreshField(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Frontend.Auth == nil {
		return ""
	}
	return fs.Manifest.Frontend.Auth.RefreshField
}
