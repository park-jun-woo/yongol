import { Global, Module } from '@nestjs/common';
import { BillingService } from './billing.service';

@Global()
@Module({
  providers: [BillingService],
  exports: [BillingService],
})
export class BillingModule {}
