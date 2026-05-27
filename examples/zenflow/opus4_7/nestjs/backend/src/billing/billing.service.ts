import { Injectable } from '@nestjs/common';

@Injectable()
export class BillingService {
  async isZeroBalance(...args: any[]): Promise<any> {
    throw new Error('BillingService.isZeroBalance not implemented');
  }

}
