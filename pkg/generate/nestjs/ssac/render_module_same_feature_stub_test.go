//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderModuleSameFeatureStub — 서브테스트 디스패치
package ssac

import "testing"

func TestRenderModuleSameFeatureStub(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"AuthModuleRegistersAuthService", subtestTestRenderModuleSameFeatureStubAuthModuleRegistersAuthService},
		{"NoCrossFeatureCallNoStub", subtestTestRenderModuleSameFeatureStubNoCrossFeatureCallNoStub},
		{"MixedSameAndCrossFeature", subtestTestRenderModuleSameFeatureStubMixedSameAndCrossFeature},
		{"NoSameFeatureCallNoStub", subtestTestRenderModuleSameFeatureStubNoSameFeatureCallNoStub},
		{"ScheduleModuleRegistersScheduleService", subtestTestRenderModuleSameFeatureStubScheduleModuleRegistersScheduleService},
	} {
		t.Run(st.name, st.fn)
	}
}
