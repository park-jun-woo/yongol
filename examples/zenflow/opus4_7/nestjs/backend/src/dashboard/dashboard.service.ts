import { Injectable } from '@nestjs/common';

@Injectable()
export class DashboardService {
  async buildExecutionDetail(...args: any[]): Promise<any> {
    throw new Error('DashboardService.buildExecutionDetail not implemented');
  }

  async summarize(...args: any[]): Promise<any> {
    throw new Error('DashboardService.summarize not implemented');
  }

}
