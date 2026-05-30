//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func httpServicePlan() *ir.ServicePlan {
	return &ir.ServicePlan{
		OperationID:     "CreateCourse",
		TriggerKind:     ir.TriggerHTTP,
		HTTPMethod:      "POST",
		URLPath:         "/courses",
		Feature:         "course",
		UsesTransaction: true,
		PathParams:      []string{"id"},
		QueryParams:     []ir.QueryParamMeta{{Name: "limit"}},
		BodyFields:      []ir.BodyFieldMeta{{Name: "title"}},
		Ops: []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{VarName: "course", Model: "Course"}},
			{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "course.created"}},
			{Kind: ir.OpAuth, Auth: &ir.AuthOp{Action: "create", Resource: "course", Message: "denied", StatusCode: 403}},
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}},
		},
	}
}

func subscribeServicePlan() *ir.ServicePlan {
	return &ir.ServicePlan{
		OperationID: "OnCourseCreated",
		TriggerKind: ir.TriggerSubscribe,
		Topic:       "course.created",
		Feature:     "course",
		Ops: []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{VarName: "log", Model: "Log"}},
		},
	}
}

func TestRenderService_HTTP(t *testing.T) {
	out, err := RenderService(httpServicePlan(), nil)
	if err != nil {
		t.Fatalf("RenderService: %v", err)
	}
	for _, want := range []string{
		"@Injectable()",
		"export class CreateCourseService {",
		"private readonly prisma: PrismaService,",
		"private readonly queue: QueueService,",
		"private readonly authz: AuthzService,",
		"billingService",
		"$transaction(async (tx)",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderService HTTP missing %q\n%s", want, out)
		}
	}
}

func TestRenderService_Subscribe(t *testing.T) {
	out, err := RenderService(subscribeServicePlan(), nil)
	if err != nil {
		t.Fatalf("RenderService: %v", err)
	}
	if !strings.Contains(out, "const message = payload;") {
		t.Errorf("subscribe alias missing\n%s", out)
	}
}

func TestRenderService_Nil(t *testing.T) {
	if _, err := RenderService(nil, nil); err == nil {
		t.Error("nil plan should error")
	}
}

func TestRenderController_HTTP(t *testing.T) {
	out, err := RenderController(httpServicePlan())
	if err != nil {
		t.Fatalf("RenderController: %v", err)
	}
	for _, want := range []string{
		"@Controller('courses')",
		"export class CreateCourseController {",
		"from '@nestjs/common';",
		"Param,",
		"Body,",
		"Query,",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderController HTTP missing %q\n%s", want, out)
		}
	}
}

func TestRenderController_Subscribe(t *testing.T) {
	out, err := RenderController(subscribeServicePlan())
	if err != nil {
		t.Fatalf("RenderController: %v", err)
	}
	if !strings.Contains(out, "Subscribe handler for topic: course.created") {
		t.Errorf("subscribe handler missing\n%s", out)
	}
	if !strings.Contains(out, "async handleOnCourseCreated(payload: any)") {
		t.Errorf("handler signature missing\n%s", out)
	}
}

func TestRenderController_Nil(t *testing.T) {
	if _, err := RenderController(nil); err == nil {
		t.Error("nil plan should error")
	}
}
