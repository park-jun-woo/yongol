//ff:type feature=stml-gen type=model topic=import-collect
//ff:what 페이지에 필요한 고유 임포트를 수집하는 구조체
package stml

// importSet collects unique imports for a page.
type importSet struct {
	react          bool
	useQuery       bool
	useMutation    bool
	useQueryClient bool
	useParams      bool
	useForm        bool
	useState       bool
	components     []string // unique component names
	customFile     string   // non-empty if custom.ts exists
}
