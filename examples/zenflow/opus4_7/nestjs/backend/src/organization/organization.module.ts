import { Module } from '@nestjs/common';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthzModule } from '../authz/authz.module';
import { GeocodingModule } from '../geocoding/geocoding.module';
import { VerifyOrgAddressController } from './verifyOrgAddress.controller';
import { VerifyOrgAddressService } from './verifyOrgAddress.service';

@Module({
  imports: [
    PrismaModule,
    AuthzModule,
    GeocodingModule,
  ],
  controllers: [
    VerifyOrgAddressController,
  ],
  providers: [
    VerifyOrgAddressService,
  ],
})
export class OrganizationModule {}
