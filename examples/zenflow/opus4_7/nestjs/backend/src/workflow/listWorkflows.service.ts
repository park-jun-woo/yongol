import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class ListWorkflowsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listWorkflows(user?: any): Promise<any> {
    await this.authz.check({
      action: 'ListWorkflows',
      resource: 'workflow',
    });
    const workflows = await this.prisma.workflow.findMany({ where: { org_id: user.org_id } });
    return {
      workflows: workflows,
    };
  }
}
