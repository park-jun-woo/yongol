import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class ListWebhooksService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listWebhooks(user?: any): Promise<any> {
    const owner = await tx.webhooks.findUnique({
      where: { id: params.id },
      select: { org_id: true },
    });
    await this.authz.check({
      action: 'ListWebhooks',
      resource: 'webhook',
      resourceId: String(params.id),
      owners: { webhooks: { org_id: owner?.org_id } },
    });
    const webhooks = await this.prisma.webhook.findMany({ where: { org_id: user.org_id } });
    return {
      webhooks: webhooks,
    };
  }
}
