#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PostsRequest {
    #[prost(string, repeated, tag = "1")]
    pub tags: ::std::vec::Vec<std::string::String>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PostsResponse {
    #[prost(message, repeated, tag = "1")]
    pub posts: ::std::vec::Vec<posts_response::Post>,
}
pub mod posts_response {
    #[derive(Clone, PartialEq, ::prost::Message, serde::Serialize, serde::Deserialize)]
    pub struct Post {
        #[prost(int32, tag = "1")]
        pub id: i32,
        #[prost(string, repeated, tag = "2")]
        pub tags: ::std::vec::Vec<std::string::String>,
        #[prost(string, tag = "3")]
        pub file_url: std::string::String,
        #[prost(int32, tag = "4")]
        pub created_at: i32,
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Rating {
    All = 0,
    Safe = 1,
    Questionable = 2,
    Explicit = 3,
}
#[doc = r" Generated server implementations."]
pub mod gallery_server {
    #![allow(unused_variables, dead_code, missing_docs)]
    use tonic::codegen::*;
    #[doc = "Generated trait containing gRPC methods that should be implemented for use with GalleryServer."]
    #[async_trait]
    pub trait Gallery: Send + Sync + 'static {
        async fn posts(
            &self,
            request: tonic::Request<super::PostsRequest>,
        ) -> Result<tonic::Response<super::PostsResponse>, tonic::Status>;
    }
    #[derive(Debug)]
    pub struct GalleryServer<T: Gallery> {
        inner: _Inner<T>,
    }
    struct _Inner<T>(Arc<T>, Option<tonic::Interceptor>);
    impl<T: Gallery> GalleryServer<T> {
        pub fn new(inner: T) -> Self {
            let inner = Arc::new(inner);
            let inner = _Inner(inner, None);
            Self { inner }
        }
        pub fn with_interceptor(inner: T, interceptor: impl Into<tonic::Interceptor>) -> Self {
            let inner = Arc::new(inner);
            let inner = _Inner(inner, Some(interceptor.into()));
            Self { inner }
        }
    }
    impl<T, B> Service<http::Request<B>> for GalleryServer<T>
    where
        T: Gallery,
        B: HttpBody + Send + Sync + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = Never;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(&mut self, _cx: &mut Context<'_>) -> Poll<Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/konachan.v1.Gallery/Posts" => {
                    #[allow(non_camel_case_types)]
                    struct PostsSvc<T: Gallery>(pub Arc<T>);
                    impl<T: Gallery> tonic::server::UnaryService<super::PostsRequest> for PostsSvc<T> {
                        type Response = super::PostsResponse;
                        type Future = BoxFuture<tonic::Response<Self::Response>, tonic::Status>;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::PostsRequest>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { (*inner).posts(request).await };
                            Box::pin(fut)
                        }
                    }
                    let inner = self.inner.clone();
                    let fut = async move {
                        let interceptor = inner.1.clone();
                        let inner = inner.0;
                        let method = PostsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = if let Some(interceptor) = interceptor {
                            tonic::server::Grpc::with_interceptor(codec, interceptor)
                        } else {
                            tonic::server::Grpc::new(codec)
                        };
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => Box::pin(async move {
                    Ok(http::Response::builder()
                        .status(200)
                        .header("grpc-status", "12")
                        .body(tonic::body::BoxBody::empty())
                        .unwrap())
                }),
            }
        }
    }
    impl<T: Gallery> Clone for GalleryServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self { inner }
        }
    }
    impl<T: Gallery> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(self.0.clone(), self.1.clone())
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: Gallery> tonic::transport::NamedService for GalleryServer<T> {
        const NAME: &'static str = "konachan.v1.Gallery";
    }
}
