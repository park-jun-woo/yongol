import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class ListWorkflowVersionsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listWorkflowVersions(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'ListWorkflowVersions',
      resource: 'workflow',
      ResourceID: params.id,
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const workflows = await this.prisma.workflow.findMany({ where: { org_id: user.org_id, root_id: wf.id } });
    return {
      workflows: workflows,
    };
  }
}
