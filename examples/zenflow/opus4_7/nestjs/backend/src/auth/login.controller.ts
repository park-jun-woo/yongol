import {
  Controller,
  Post,
  Body,
  Req,
} from '@nestjs/common';
import { LoginService } from './login.service';

@Controller('auth')
export class LoginController {
  constructor(private readonly service: LoginService) {}

  @Post('login')
  async login(
    @Req() req: any,
    @Body() body: any,
  ) {
    return this.service.login(body, req.user);
  }
}
