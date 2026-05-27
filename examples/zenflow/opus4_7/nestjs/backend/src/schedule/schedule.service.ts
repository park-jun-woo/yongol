import { Injectable } from '@nestjs/common';

@Injectable()
export class ScheduleService {
  async buildKey(...args: any[]): Promise<any> {
    throw new Error('ScheduleService.buildKey not implemented');
  }

  async parseCron(...args: any[]): Promise<any> {
    throw new Error('ScheduleService.parseCron not implemented');
  }

}
