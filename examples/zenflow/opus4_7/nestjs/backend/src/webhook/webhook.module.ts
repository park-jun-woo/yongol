import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { WebhookdeliveryModule } from '../webhookdelivery/webhookdelivery.module';
import { CreateWebhookController } from './createWebhook.controller';
import { CreateWebhookService } from './createWebhook.service';
import { DeleteWebhookController } from './deleteWebhook.controller';
import { DeleteWebhookService } from './deleteWebhook.service';
import { ListWebhooksController } from './listWebhooks.controller';
import { ListWebhooksService } from './listWebhooks.service';
import { OnWorkflowExecutedController } from './onWorkflowExecuted.controller';
import { OnWorkflowExecutedService } from './onWorkflowExecuted.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
    WebhookdeliveryModule,
  ],
  controllers: [
    CreateWebhookController,
    DeleteWebhookController,
    ListWebhooksController,
    OnWorkflowExecutedController,
  ],
  providers: [
    CreateWebhookService,
    DeleteWebhookService,
    ListWebhooksService,
    OnWorkflowExecutedService,
  ],
})
export class WebhookModule {}
