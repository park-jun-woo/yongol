//ff:func feature=gen-gogin type=util control=sequence
//ff:what lcFirst — 첫 글자 소문자로

package ssac

func lcFirst(s string) string {
	if s == "" || s[0] < 'A' || s[0] > 'Z' {
		return s
	}
	return string(s[0]+('a'-'A')) + s[1:]
}
