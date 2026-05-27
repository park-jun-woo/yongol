import { Injectable } from '@nestjs/common';

@Injectable()
export class VersionService {
  async nextVersion(...args: any[]): Promise<any> {
    throw new Error('VersionService.nextVersion not implemented');
  }

  async resolveRootID(...args: any[]): Promise<any> {
    throw new Error('VersionService.resolveRootID not implemented');
  }

}
