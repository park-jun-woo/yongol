import { Injectable } from '@nestjs/common';

@Injectable()
export class GeocodingService {
  async geocode(...args: any[]): Promise<any> {
    throw new Error('GeocodingService.geocode not implemented');
  }

}
