#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let DISCORD_BOT_TOKEN =
        std::env::var("DISCORD_BOT_TOKEN").expect("DISCORD_BOT_TOKEN envvar required");
    let DISCORD_BOT_CHANNEL =
        std::env::var("DISCORD_BOT_CHANNEL").expect("DISCORD_BOT_CHANNEL envvar required");

    let DISCORD_API_URL = "https://discord.com/api/v10";
    let reqwest_client = reqwest::Client::new();

    let mut header_map = reqwest::header::HeaderMap::new();
    header_map.insert(
        "Authorization",
        reqwest::header::HeaderValue::from_str(format!("Bot {}", DISCORD_BOT_TOKEN).as_str())
            .unwrap(),
    );

    let mut content_map = std::collections::HashMap::new();
    content_map.insert("content", "Test rust");

    let resp = reqwest_client
        .post(format!(
            "{}/channels/{}/messages",
            DISCORD_API_URL, DISCORD_BOT_CHANNEL
        ))
        .headers(header_map)
        .json(&content_map)
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    println!("{:#?}", resp);
    Ok(())
}
