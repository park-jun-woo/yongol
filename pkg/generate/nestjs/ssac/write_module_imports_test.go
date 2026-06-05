//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestWriteModuleImports — TestWriteModuleImports — NestJS module 파일 상단 import 블록 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteModuleImports(t *testing.T) {
	var b strings.Builder
	plans := []*ir.ServicePlan{{OperationID: "CreateCourse"}}
	deps := moduleDeps{
		NeedsQueue:           true,
		NeedsAuthz:           true,
		NeedsSameFeatureStub: true,
		CrossFeatures:        []string{"billing"},
	}
	writeModuleImports(&b, "Course", plans, deps, "CourseStubService")

	out := b.String()
	for _, want := range []string{
		"import { Module } from '@nestjs/common';\n",
		"import { PrismaModule } from '../prisma/prisma.module';\n",
		"import { QueueModule } from '../queue/queue.module';\n",
		"import { AuthzModule } from '../authz/authz.module';\n",
		"import { BillingModule } from '../billing/billing.module';\n",
		"import { CreateCourseController } from './createCourse.controller';\n",
		"import { CreateCourseService } from './createCourse.service';\n",
		"import { CourseStubService } from './course.service';\n",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("imports missing %q\n--- got ---\n%s", want, out)
		}
	}
}
