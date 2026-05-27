import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { LoginController } from './login.controller';
import { LoginService } from './login.service';

@Module({
  imports: [
    PrismaModule,
  ],
  controllers: [
    LoginController,
  ],
  providers: [
    LoginService,
  ],
})
export class AuthModule {}
