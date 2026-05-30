import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { WorkflowService } from '../../workflow/workflow.service';

@Injectable()
export class AutoAssignWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly workflowService: WorkflowService,
  ) {}

  async autoAssignWorkflow(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'AutoAssignWorkflow',
        resource: 'workflow',
        ResourceID: params.id,
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      const memberCount = await tx.user.findUnique({ where: { org_id: wf.org_id } });
      const match = await this.workflowService.matchMember(memberCount, wf.trigger_event);
      await tx.workflow.update({ where: { id: wf.id }, data: { confidence: match.confidence, member_id: match.member_id } });
      const updated = await tx.workflow.findUnique({ where: { id: wf.id } });
      if (!updated) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      return {
        workflow: updated,
      };
    });
  }
}
