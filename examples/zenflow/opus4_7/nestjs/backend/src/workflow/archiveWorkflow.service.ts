import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class ArchiveWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async archiveWorkflow(params: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.workflow.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'ArchiveWorkflow',
        resource: 'workflow',
        resourceId: String(params.id),
        owners: { workflow: { org_id: owner?.org_id } },
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      // @state workflows.ArchiveWorkflow — transition guard
      const allowed_ArchiveWorkflow: Record<string, boolean> = {
        'active': true,
      };
      if (!allowed_ArchiveWorkflow[wf.status]) {
        throw new HttpException('Cannot archive workflow', HttpStatus.CONFLICT);
      }
      await tx.workflow.update({ where: { id: wf.id }, data: { status: 'archived' } });
      const updated = await tx.workflow.findUnique({ where: { id: wf.id } });
      return {
        workflow: updated,
      };
    });
  }
}
