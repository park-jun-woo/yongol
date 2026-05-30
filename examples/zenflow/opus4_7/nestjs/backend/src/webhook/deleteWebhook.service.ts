import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthzService } from '../../authz/authz.service';

@Injectable()
export class DeleteWebhookService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
  ) {}

  async deleteWebhook(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      await this.authz.check({
        action: 'DeleteWebhook',
        resource: 'webhook',
        ResourceID: params.id,
      });
      const webhook = await tx.webhook.findUnique({ where: { id: params.id } });
      if (!webhook) {
        throw new HttpException('Webhook not found', HttpStatus.NOT_FOUND);
      }
      await tx.webhook.delete({ where: { id: webhook.id } });
    });
  }
}
