//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestCollectFieldArgs — 서브테스트 디스패치
package ssac

import "testing"

func TestCollectFieldArgs(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"GetWithArgsAndPagination", subtestTestCollectFieldArgsGetWithArgsAndPagination},
		{"GetNil", subtestTestCollectFieldArgsGetNil},
		{"Post", subtestTestCollectFieldArgsPost},
		{"PostNil", subtestTestCollectFieldArgsPostNil},
		{"Put", subtestTestCollectFieldArgsPut},
		{"PutNil", subtestTestCollectFieldArgsPutNil},
		{"Delete", subtestTestCollectFieldArgsDelete},
		{"DeleteNil", subtestTestCollectFieldArgsDeleteNil},
		{"Auth", subtestTestCollectFieldArgsAuth},
		{"AuthNil", subtestTestCollectFieldArgsAuthNil},
		{"State", subtestTestCollectFieldArgsState},
		{"StateNil", subtestTestCollectFieldArgsStateNil},
		{"Call", subtestTestCollectFieldArgsCall},
		{"CallNil", subtestTestCollectFieldArgsCallNil},
		{"Eval", subtestTestCollectFieldArgsEval},
		{"EvalNil", subtestTestCollectFieldArgsEvalNil},
		{"PublishPayloadAndOptions", subtestTestCollectFieldArgsPublishPayloadAndOptions},
		{"PublishNil", subtestTestCollectFieldArgsPublishNil},
		{"DefaultUnknownKind", subtestTestCollectFieldArgsDefaultUnknownKind},
	} {
		t.Run(st.name, st.fn)
	}
}
