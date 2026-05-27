import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class CreateWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async createWorkflow(body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'CreateWorkflow',
        resource: 'workflow',
      });
      const wf = await tx.workflow.create({ data: { org_id: user.org_id, title: body.title, trigger_event: body.trigger_event } });
      return {
        workflow: wf,
      };
    });
  }
}
