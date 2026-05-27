import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';
import { VersionService } from '../../version/version.service';

@Injectable()
export class CreateWorkflowVersionService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly versionService: VersionService,
  ) {}

  async createWorkflowVersion(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'CreateWorkflowVersion',
        resource: 'workflow',
        ResourceID: params.id,
      });
      const wf = await tx.workflow.findUnique({ where: { id: params.id } });
      if (!wf) {
        throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
      }
      const rootResult = await this.versionService.resolveRootID(wf.root_workflow_id, wf.id);
      const versionResult = await this.versionService.nextVersion(wf.version);
      const newWf = await tx.workflow.create({ data: { org_id: wf.org_id, root_workflow_id: rootResult.root_id, title: wf.title, trigger_event: wf.trigger_event, version: versionResult.version } });
      await tx.action.update({ where: { new_workflow_id: newWf.id }, data: { source_workflow_id: wf.id } });
      return {
        workflow: newWf,
      };
    });
  }
}
