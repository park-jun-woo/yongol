import { Injectable } from '@nestjs/common';

@Injectable()
export class SessionService {
  async delete(...args: any[]): Promise<any> {
    throw new Error('SessionService.delete not implemented');
  }

  async get(...args: any[]): Promise<any> {
    throw new Error('SessionService.get not implemented');
  }

  async set(...args: any[]): Promise<any> {
    throw new Error('SessionService.set not implemented');
  }

}
