import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';

@Injectable()
export class CreateWebhookService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async createWebhook(body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.webhooks.findUnique({
        where: { id: params.id },
        select: { org_id: true },
      });
      await this.authz.check({
        action: 'CreateWebhook',
        resource: 'webhook',
        resourceId: String(params.id),
        owners: { webhooks: { org_id: owner?.org_id } },
      });
      const webhook = await tx.webhook.create({ data: { event_type: body.event_type, org_id: user.org_id, url: body.url } });
      return {
        webhook: webhook,
      };
    });
  }
}
