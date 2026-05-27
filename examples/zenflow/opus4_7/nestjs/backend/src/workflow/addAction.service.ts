import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class AddActionService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async addAction(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.workflows.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'AddAction',
        resource: 'workflow',
        ResourceID: params.id,
        resourceId: String(params.id),
        owners: { workflows: { org_id: owner?.org_id } },
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      const action = await tx.action.create({ data: { action_type: body.action_type, config: body.config, sequence_order: body.sequence_order, workflow_id: wf.workflow_id } });
      return {
        action: action,
      };
    });
  }
}
