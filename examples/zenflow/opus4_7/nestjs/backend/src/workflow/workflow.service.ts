import { Injectable } from '@nestjs/common';

@Injectable()
export class WorkflowService {
  async matchMember(...args: any[]): Promise<any> {
    throw new Error('WorkflowService.matchMember not implemented');
  }

}
