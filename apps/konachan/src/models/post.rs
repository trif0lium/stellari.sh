#[derive(Clone, PartialEq, serde::Serialize, serde::Deserialize)]
pub struct Post {
    pub id: i32,
    pub tags: std::string::String,
    pub file_url: std::string::String,
    pub created_at: i32,
}
