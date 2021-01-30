import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { KonachanModule } from './konachan/konachan.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: ['.env.development', '.env'],
    }),
    KonachanModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
