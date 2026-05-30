import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class CreateWebhookService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async createWebhook(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'CreateWebhook',
        resource: 'webhook',
      });
      const webhook = await tx.webhook.create({ data: { event_type: params.event_type, org_id: user.org_id, url: params.url } });
      return {
        webhook: webhook,
      };
    });
  }
}
