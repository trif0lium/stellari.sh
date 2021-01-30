/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable } from 'rxjs';

export const protobufPackage = 'konachan.v1';

export enum Rating {
  ALL = 0,
  SAFE = 1,
  QUESTIONABLE = 2,
  EXPLICIT = 3,
  UNRECOGNIZED = -1,
}

export interface PostsRequest {
  tags: string[];
}

export interface PostsResponse {
  posts: PostsResponse_Post[];
}

export interface PostsResponse_Post {
  id: number;
  tags: string[];
  fileUrl: string;
  createdAt: number;
}

export const KONACHAN_V1_PACKAGE_NAME = 'konachan.v1';

export interface GalleryClient {
  posts(request: PostsRequest): Observable<PostsResponse>;
}

export interface GalleryController {
  posts(
    request: PostsRequest,
  ): Promise<PostsResponse> | Observable<PostsResponse> | PostsResponse;
}

export function GalleryControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods = ['posts'];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcMethod('Gallery', method)(
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
      GrpcStreamMethod('Gallery', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
  };
}

export const GALLERY_SERVICE_NAME = 'Gallery';
