import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { WebhookdeliveryService } from '../../webhookdelivery/webhookdelivery.service';

@Injectable()
export class OnWorkflowExecutedService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly webhookdeliveryService: WebhookdeliveryService,
  ) {}

  async onWorkflowExecuted(payload: any): Promise<any> {
    await this.webhookdeliveryService.deliver(message.status, 'simulated');
  }
}
