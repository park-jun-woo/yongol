import { Module } from '@nestjs/common';
import { PrismaModule } from '../../prisma/prisma.module';
import { VerifyOrgAddressController } from './verifyOrgAddress.controller';
import { VerifyOrgAddressService } from './verifyOrgAddress.service';

@Module({
  imports: [
    PrismaModule,
  ],
  controllers: [
    VerifyOrgAddressController,
  ],
  providers: [
    VerifyOrgAddressService,
  ],
})
export class OrganizationModule {}
