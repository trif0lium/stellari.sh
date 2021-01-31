import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { KONACHAN_V1_PACKAGE_NAME } from 'src/generated/apis/konachan';
import { KonachanController } from './konachan.controller';
import { KonachanService } from './konachan.service';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: KONACHAN_V1_PACKAGE_NAME,
        imports: [ConfigModule],
        inject: [ConfigService],
        useFactory: async (configService: ConfigService<NodeJS.ProcessEnv>) => ({
          name: KONACHAN_V1_PACKAGE_NAME,
          transport: Transport.GRPC,
          options: {
            package: KONACHAN_V1_PACKAGE_NAME,
            protoPath: join(__dirname, '../../apis/konachan.proto'),
            url: configService.get<string>('SVC_KONACHAN_GRPC_URL'),
          },
        }),
      },
    ]),
  ],
  controllers: [KonachanController],
  providers: [KonachanService],
})
export class KonachanModule {}
