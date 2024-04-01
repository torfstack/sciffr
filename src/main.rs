use clap::Parser;
use crate::cipher::CipherTrait;

mod config;
mod cipher;

/// Encrypt and decrypt configuration values without affecting keys
#[derive(Parser, Debug)]
#[command(version, about)]
struct Params {
    /// Path to the configuration file
    #[arg(short, long)]
    file: String,
}

fn main() {
    let args = Params::parse();

    let config = config::read_in_config();
    let cipher = cipher::Cipher::new(config.encryption.key);

    let data = std::fs::read_to_string(&args.file).unwrap();
    let encrypted = cipher.encrypt(&data);
    println!("Encrypted: {}", encrypted);
    let decrypted = cipher.decrypt(&encrypted);
    println!("Decrypted: {}", decrypted);
}
