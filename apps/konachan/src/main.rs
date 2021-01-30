mod apis;
mod models;

use apis::konachan::v1::gallery_server;
use apis::konachan::v1::posts_response::Post;
use apis::konachan::v1::{PostsRequest, PostsResponse};

#[derive(Debug, Default)]
pub struct GalleryImpl {}

#[tonic::async_trait]
impl gallery_server::Gallery for GalleryImpl {
    async fn posts(
        &self,
        request: tonic::Request<PostsRequest>,
    ) -> Result<tonic::Response<PostsResponse>, tonic::Status> {
        let req = request.into_inner();
        let tags = req.tags;

        if tags.len() == 0 {
            return Err(tonic::Status::failed_precondition(
                "Tags should not be empty.",
            ));
        }

        let mut url = reqwest::Url::parse("https://konachan.com/post.json?limit=100").unwrap();
        url.query_pairs_mut()
            .append_pair("tags", &(tags.clone().join(" ")));
        let mut ret = PostsResponse { posts: [].to_vec() };
        let resp = reqwest::blocking::get(url);
        let _ = match resp {
            Ok(res) => {
                let json = res.json::<Vec<models::post::Post>>();
                let _ = match json {
                    Ok(posts) => {
                        ret.posts = posts
                            .into_iter()
                            .map(|post| {
                                return Post {
                                    id: post.id,
                                    tags: [].to_vec(),
                                    file_url: post.file_url,
                                    created_at: post.created_at,
                                };
                            })
                            .collect();
                    }
                    Err(error) => {
                        println!("{:#?}", error);
                    }
                };
            }
            Err(error) => {
                println!("{:#?}", error);
            }
        };

        Ok(tonic::Response::new(ret))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:50051".parse()?;
    let server = GalleryImpl::default();
    tonic::transport::Server::builder()
        .add_service(gallery_server::GalleryServer::new(server))
        .serve(addr)
        .await?;
    Ok(())
}
