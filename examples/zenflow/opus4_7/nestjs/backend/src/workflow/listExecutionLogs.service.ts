import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class ListExecutionLogsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listExecutionLogs(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'ListExecutionLogs',
      resource: 'workflow',
      ResourceID: params.id,
    });
    const wf = await this.prisma.workflow.findUnique({ where: { id: params.id } });
    if (!wf) {
      throw new HttpException('Workflow not found', HttpStatus.NOT_FOUND);
    }
    const execution_logs = await this.prisma.executionLog.findMany({ where: { workflow_id: wf.id } });
    return {
      execution_logs: execution_logs,
    };
  }
}
