import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { KonachanModule } from './konachan/konachan.module';
import { resolve } from 'path';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: [resolve(__dirname, '../.env')],
    }),
    KonachanModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
