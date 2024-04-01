use crate::cipher::CipherTrait;

mod config;
mod cipher;

fn main() {
    let config = config::read_in_config();
    let cipher = cipher::Cipher::new(config.encryption.key);
    let encrypted = cipher.encrypt("Hello, world!");
    println!("Encrypted: {}", encrypted);
    let decrypted = cipher.decrypt(&encrypted);
    println!("Decrypted: {}", decrypted);
}
