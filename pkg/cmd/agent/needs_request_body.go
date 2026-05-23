//ff:func feature=agent type=helper control=sequence
//ff:what needsRequestBody — HTTP method가 request body를 필요로 하는지 판별

package agent

func needsRequestBody(method string) bool {
	return method == "post" || method == "put" || method == "patch"
}
