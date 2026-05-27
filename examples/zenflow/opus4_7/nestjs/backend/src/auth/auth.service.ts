import { Injectable } from '@nestjs/common';

@Injectable()
export class AuthService {
  async issueToken(...args: any[]): Promise<any> {
    throw new Error('AuthService.issueToken not implemented');
  }

}
