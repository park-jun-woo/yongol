import { Injectable } from '@nestjs/common';

@Injectable()
export class WebhookdeliveryService {
  async deliver(...args: any[]): Promise<any> {
    throw new Error('WebhookdeliveryService.deliver not implemented');
  }

}
