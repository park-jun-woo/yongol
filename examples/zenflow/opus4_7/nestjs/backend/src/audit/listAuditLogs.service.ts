import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class ListAuditLogsService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listAuditLogs(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'ListAuditLogs',
      resource: 'audit_log',
    });
    const items = await this.prisma.auditLog.findMany({ where: { filter_action: params.action, org_id: user.org_id, page_offset: params.page, per_page: params.per_page } });
    const total = await this.prisma.auditLog.findUnique({ where: { filter_action: params.action, org_id: user.org_id } });
    return {
      items: items,
      total: total,
    };
  }
}
