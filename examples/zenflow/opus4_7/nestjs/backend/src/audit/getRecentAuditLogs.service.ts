import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class GetRecentAuditLogsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async getRecentAuditLogs(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'GetRecentAuditLogs',
      resource: 'audit_log',
    });
    const items = await this.prisma.auditLog.findMany({ where: { filter_action: params, org_id: user.org_id, page_offset: 0, per_page: 10 } });
    return {
      items: items,
    };
  }
}
