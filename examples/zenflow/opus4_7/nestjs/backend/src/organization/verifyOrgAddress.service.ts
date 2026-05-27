import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { AuthzService } from '../authz/authz.service';
import { GeocodingService } from '../geocoding/geocoding.service';

@Injectable()
export class VerifyOrgAddressService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly authz: AuthzService,
    private readonly geocodingService: GeocodingService,
  ) {}

  async verifyOrgAddress(params: any, body: any, user?: any): Promise<any> {
    return this.prisma.$transaction(async (tx) => {
      const owner = await tx.organizations.findUnique({
        where: { id: params.id },
        select: { id: true },
      });
      await this.authz.check({
        action: 'VerifyOrgAddress',
        resource: 'organization',
        ResourceID: params.id,
        resourceId: String(params.id),
        owners: { organizations: { id: owner?.id } },
      });
      const org = await tx.organization.findUnique({ where: { id: params.id } });
      if (!org) {
        throw new HttpException('Organization not found', HttpStatus.NOT_FOUND);
      }
      const geo = await this.geocodingService.geocode(body.address);
      await tx.organization.update({ where: { id: org.id }, data: { address_verified: geo.address_verified, latitude: geo.latitude, longitude: geo.longitude } });
      const updated_org = await tx.organization.findUnique({ where: { id: org.id } });
      if (!updated_org) {
        throw new HttpException('Organization not found', HttpStatus.NOT_FOUND);
      }
      return {
        organization: updated_org,
      };
    });
  }
}
