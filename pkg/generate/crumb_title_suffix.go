//ff:func feature=generate type=util control=sequence
//ff:what crumbTitleSuffix — 동적 crumb 라벨의 document.title 결합 꼬리(" · <앱명>") 산출 (앱명 없으면 빈 문자열)

package generate

import "github.com/park-jun-woo/yongol/pkg/yongol"

// crumbTitleSuffix returns the " · <app name>" tail the Phase006 dynamic
// crumb effect appends when it updates document.title — the same join
// addDocumentTitle uses for the static mount title, so the dynamic and
// static titles always share one format. "" without a manifest app name
// (the dynamic title is then the bare label, matching the static rule).
func crumbTitleSuffix(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Metadata.Name == "" {
		return ""
	}
	return " · " + fs.Manifest.Metadata.Name
}
