import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthService } from '../auth/auth.service';

@Injectable()
export class LoginService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authService: AuthService,
  ) {}

  async login(body: any, user?: any): Promise<any> {
    const user_result = await this.prisma.user.findUnique({ where: { email: body.email } });
    if (!user_result) {
      throw new HttpException('Invalid credentials', HttpStatus.UNAUTHORIZED);
    }
    // TODO: bcrypt.compare(body.password, user_result.password_hash)
    const token = await this.authService.issueToken(user.email, user.id, user.org_id, user.role);
    return {
      access_token: token.AccessToken,
    };
  }
}
