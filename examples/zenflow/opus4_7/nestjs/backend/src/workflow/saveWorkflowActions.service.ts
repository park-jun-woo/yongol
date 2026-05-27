import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class SaveWorkflowActionsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async saveWorkflowActions(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.workflows.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'SaveWorkflowActions',
        resource: 'workflow',
        ResourceID: params.id,
        resourceId: String(params.id),
        owners: { workflows: { org_id: owner?.org_id } },
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      await tx.action.deleteMany({ where: { workflow_id: wf.workflow_id } });
      await tx.action.update({ where: { id: params.id }, data: { items: body.actions_json, workflow_id: wf.workflow_id } });
      return {
        message: "ok",
      };
    });
  }
}
