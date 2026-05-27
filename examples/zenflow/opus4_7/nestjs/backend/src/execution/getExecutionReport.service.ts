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
    await this.authz.check({
      action: 'GetExecutionReport',
      resource: 'execution_log',
    });
    const log = await this.prisma.executionLog.findUnique({ where: { id: params.id } });
    if (!log) {
      throw new HttpException('Execution log not found', HttpStatus.NOT_FOUND);
    }
    return {
      report_key: log.reportKey,
    };
  }
}
