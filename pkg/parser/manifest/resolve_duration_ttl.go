//ff:func feature=manifest type=util control=sequence
//ff:what resolveDurationTTL — Auth duration 문자열을 초 단위 int64 로 변환

package manifest

import "time"

// resolveDurationTTL extracts a duration string from the Auth config via the
// given accessor, parses it with time.ParseDuration, and converts to seconds.
func resolveDurationTTL(auth *Auth, accessor func(*Auth) string) (RefValue, bool) {
	if auth == nil {
		return RefValue{}, false
	}
	raw := accessor(auth)
	if raw == "" {
		return RefValue{}, false
	}
	d, err := time.ParseDuration(raw)
	if err != nil {
		return RefValue{}, false
	}
	seconds := int64(d.Seconds())
	return RefValue{
		Raw:    raw,
		GoLit:  intToStr(seconds),
		GoType: "int64",
	}, true
}
