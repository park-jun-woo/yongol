import { Global, Module } from '@nestjs/common';
import { VersionService } from './version.service';

@Global()
@Module({
  providers: [VersionService],
  exports: [VersionService],
})
export class VersionModule {}
