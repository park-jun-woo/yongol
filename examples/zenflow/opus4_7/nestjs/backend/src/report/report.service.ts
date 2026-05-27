import { Injectable } from '@nestjs/common';

@Injectable()
export class ReportService {
  async generateReport(...args: any[]): Promise<any> {
    throw new Error('ReportService.generateReport not implemented');
  }

}
