/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable } from 'rxjs';

export const protobufPackage = 'nhentai.v1';

export interface GalleryRequest {
  id: number;
}

export interface GalleryResponse {
  id: number;
  mediaId: number;
  title: GalleryResponse_Title | undefined;
}

export interface GalleryResponse_Title {
  japanese: string;
  english: string;
  pretty: string;
}

export const NHENTAI_V1_PACKAGE_NAME = 'nhentai.v1';

export interface NHClient {
  gallery(request: GalleryRequest): Observable<GalleryResponse>;
}

export interface NHController {
  gallery(
    request: GalleryRequest,
  ): Promise<GalleryResponse> | Observable<GalleryResponse> | GalleryResponse;
}

export function NHControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods = ['gallery'];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcMethod('NH', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
    const grpcStreamMethods = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcStreamMethod('NH', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
  };
}

export const N_H_SERVICE_NAME = 'NH';
