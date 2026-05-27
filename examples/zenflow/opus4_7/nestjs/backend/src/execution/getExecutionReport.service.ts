import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class GetExecutionReportService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async getExecutionReport(params: any, user?: any): Promise<any> {
    const owner = await tx.execution_logs.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'GetExecutionReport',
      resource: 'execution_log',
      resourceId: String(params.id),
      owners: { execution_logs: { org_id: owner?.org_id } },
    });
    const log = await this.prisma.executionLog.findUnique({ where: { id: params.id } });
    if (!log) {
      throw new HttpException('Execution log not found', HttpStatus.NOT_FOUND);
    }
    return {
      report_key: log.ReportKey,
    };
  }
}
