import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class ListWebhooksService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async listWebhooks(params: any, body: any, user?: any): Promise<any> {
    await this.authz.check({
      action: 'ListWebhooks',
      resource: 'webhook',
    });
    const webhooks = await this.prisma.webhook.findMany({ where: { org_id: user.org_id } });
    return {
      webhooks: webhooks,
    };
  }
}
