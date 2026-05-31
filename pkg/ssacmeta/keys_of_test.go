//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestloadPackageInterfaceEntry — loadPackageInterfaceEntry() dir/파일/키 폴백 분기
package ssacmeta

func keysOf(m map[string]*PackageInterface) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
