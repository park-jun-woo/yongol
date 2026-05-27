import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class PauseWorkflowService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async pauseWorkflow(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'PauseWorkflow',
        resource: 'workflow',
        ResourceID: params.id,
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      // @state workflows.PauseWorkflow — transition guard
      const allowed_PauseWorkflow: Record<string, string[]> = {
        // TODO: populate from Mermaid stateDiagram 'workflows'
      };
      if (!(allowed_PauseWorkflow[wf.status] || []).includes('PauseWorkflow')) {
        throw new HttpException('Cannot pause workflow', HttpStatus.CONFLICT);
      }
      await tx.workflow.update({ where: { id: wf.id }, data: { status: 'paused' } });
      const updated = await tx.workflow.findUnique({ where: { id: wf.id } });
      return {
        workflow: updated,
      };
    });
  }
}
