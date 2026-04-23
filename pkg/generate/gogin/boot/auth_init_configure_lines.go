//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitConfigureLines — auth.Configure(auth.Config{...}) 블록 라인 생성

package boot

import "fmt"

// authInitConfigureLines returns the auth.Configure(...) call wrapping the
// dynamic SecretEnv / AccessName / RefreshName values together with the
// static TTL + SameSite references assigned earlier in the block.
func authInitConfigureLines(cfg authInitConfig) []string {
	return []string{
		`auth.Configure(auth.Config{`,
		fmt.Sprintf("	SecretEnv:  %q,", cfg.SecretEnv),
		`	AccessTTL:  accessTTL,`,
		`	RefreshTTL: refreshTTL,`,
		`	Mode:       authMode,`,
		`	CookieAttrs: auth.CookieAttrs{`,
		fmt.Sprintf("		AccessName:  %q,", cfg.AccessName),
		fmt.Sprintf("		RefreshName: %q,", cfg.RefreshName),
		`		SameSite:    sameSite,`,
		`		AccessTTL:   accessTTL,`,
		`		RefreshTTL:  refreshTTL,`,
		`	},`,
		`})`,
	}
}
