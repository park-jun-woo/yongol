import { Injectable } from '@nestjs/common';

@Injectable()
export class WorkerService {
  async processAction(...args: any[]): Promise<any> {
    throw new Error('WorkerService.processAction not implemented');
  }

}
