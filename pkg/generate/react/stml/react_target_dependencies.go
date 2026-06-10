//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 페이지 목록에서 필요한 npm 의존성 목록을 수집한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func (r *ReactTarget) Dependencies(pages []stmlparser.PageSpec) map[string]string {
	deps := map[string]string{}
	for _, page := range pages {
		is := collectImports(page, "")
		if is.useQuery || is.useMutation || is.useQueryClient {
			deps["@tanstack/react-query"] = "^5"
		}
		if is.useForm {
			deps["react-hook-form"] = "^7"
		}
		if is.useParams || is.useLink {
			deps["react-router-dom"] = "^6"
		}
		if is.useZod {
			deps["zod"] = "^3"
			deps["@hookform/resolvers"] = "^3"
		}
	}
	return deps
}
