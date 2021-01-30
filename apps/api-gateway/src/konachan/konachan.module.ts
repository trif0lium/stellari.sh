import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { KONACHAN_V1_PACKAGE_NAME } from 'src/generated/apis/konachan';
import { KonachanController } from './konachan.controller';
import { KonachanService } from './konachan.service';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: KONACHAN_V1_PACKAGE_NAME,
        transport: Transport.GRPC,
        options: {
          package: KONACHAN_V1_PACKAGE_NAME,
          protoPath: join(__dirname, '../../apis/konachan.proto'),
        },
      },
    ]),
  ],
  controllers: [KonachanController],
  providers: [KonachanService],
})
export class KonachanModule {}
