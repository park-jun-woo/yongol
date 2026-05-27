//ff:func feature=gen-fastapi type=util control=sequence
//ff:what addExtPkgRef — 외부 패키지 참조를 importData 맵에 추가

package ssac

// addExtPkgRef registers an external package function reference.
func addExtPkgRef(d *importData, pkg, fn string) {
	if d.ExtPkgs[pkg] == nil {
		d.ExtPkgs[pkg] = make(map[string]bool)
	}
	d.ExtPkgs[pkg][fn] = true
}
