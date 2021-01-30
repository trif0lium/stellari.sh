fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_client(false)
        .out_dir("src/apis")
        .type_attribute(
            ".konachan.v1.PostsResponse.Post",
            "#[derive(serde::Serialize, serde::Deserialize)]",
        )
        .compile(&["apis/konachan.proto"], &["apis"])
        .unwrap();
    Ok(())
}
