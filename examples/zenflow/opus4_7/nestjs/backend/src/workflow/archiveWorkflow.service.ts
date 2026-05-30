import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class ArchiveWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async archiveWorkflow(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'ArchiveWorkflow',
        resource: 'workflow',
        ResourceID: params.id,
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      // @state workflows.ArchiveWorkflow — transition guard
      const allowed_ArchiveWorkflow: Record<string, string[]> = {
        // TODO: populate from Mermaid stateDiagram 'workflows'
      };
      if (!(allowed_ArchiveWorkflow[wf.status] || []).includes('ArchiveWorkflow')) {
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
