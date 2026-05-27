import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../../prisma/prisma.service';
import { AuthService } from '../../auth/auth.service';

@Injectable()
export class LoginService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authService: AuthService,
  ) {}

  async login(params: any, body: any, user?: any): Promise<any> {
    const user = await this.prisma.user.findUnique({ where: { email: request.email } });
    if (!user) {
      throw new HttpException('Invalid credentials', HttpStatus.UNAUTHORIZED);
    }
    // TODO: bcrypt.compare(request.password, user.password_hash)
    const token = await this.authService.issueToken(user.email, user.id, user.org_id, user.role);
    return {
      access_token: token.AccessToken,
    };
  }
}
