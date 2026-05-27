import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class ListExecutionLogsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listExecutionLogs(params: any, user?: any): Promise<any> {
    const owner = await tx.workflows.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'ListExecutionLogs',
      resource: 'workflow',
      ResourceID: params.id,
      resourceId: String(params.id),
      owners: { workflows: { org_id: owner?.org_id } },
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const execution_logs = await this.prisma.executionLog.findMany({ where: { workflow_id: wf.workflow_id } });
    return {
      execution_logs: execution_logs,
    };
  }
}
