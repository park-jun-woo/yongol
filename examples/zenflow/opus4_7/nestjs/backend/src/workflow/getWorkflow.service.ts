import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class GetWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async getWorkflow(params: any, user?: any): Promise<any> {
    const owner = await tx.workflows.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'GetWorkflow',
      resource: 'workflow',
      ResourceID: params.id,
      resourceId: String(params.id),
      owners: { workflows: { org_id: owner?.org_id } },
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    return {
      workflow: wf,
    };
  }
}
