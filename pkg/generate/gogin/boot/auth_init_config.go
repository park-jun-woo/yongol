//ff:type feature=gen-gogin type=model
//ff:what authInitConfig — blockAuthInit 가 사용하는 manifest 파생 설정 번들

package boot

// authInitConfig carries every resolved value that blockAuthInit needs to
// emit into main.go. Keeping the resolution separate from line rendering
// keeps nesting flat and each helper focused on one concern.
type authInitConfig struct {
	SecretEnv   string
	AccessTTL   string
	RefreshTTL  string
	Mode        string
	SameSite    string
	AccessName  string
	RefreshName string
	DetectReuse bool
}
