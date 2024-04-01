use clap::Parser;

mod config;
mod cipher;
mod file;

/// Encrypt and decrypt configuration values without affecting keys
#[derive(Parser, Debug)]
#[command(version, about)]
struct Params {
    /// Path to the configuration file
    #[arg(short, long)]
    file: String,

    /// Key of value in configuration file
    #[arg(short, long)]
    key: String,
}

fn main() {
    let args = Params::parse();

    let config = config::read_in_config();

    let mut file_encryptor = file::FileEncryptor::new(args.file, config.encryption.key);
    file_encryptor.encrypt(String::from(args.key));
}
