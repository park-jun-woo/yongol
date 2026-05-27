import { Global, Module } from '@nestjs/common';
import { WebhookdeliveryService } from './webhookdelivery.service';

@Global()
@Module({
  providers: [WebhookdeliveryService],
  exports: [WebhookdeliveryService],
})
export class WebhookdeliveryModule {}
