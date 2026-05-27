import {
  Controller,
  Post,
  Param,
  Body,
  Req,
} from '@nestjs/common';
import { VerifyOrgAddressService } from './verifyOrgAddress.service';

@Controller('organizations')
export class VerifyOrgAddressController {
  constructor(private readonly service: VerifyOrgAddressService) {}

  @Post(':id/verify-address')
  async verifyOrgAddress(
    @Req() req: any,
    @Param() params: any,
    @Body() body: any,
  ) {
    return this.service.verifyOrgAddress(params, body, req.user);
  }
}
