use aes;
use aes::cipher::KeyInit;
use aes_gcm::aead::{Aead, Nonce, OsRng};
use aes_gcm::{AeadCore, Aes256Gcm};
use base64::prelude::*;

pub struct Cipher {
    key: [u8; 32],
}

pub trait CipherTrait {
    fn new(key: String) -> Self;
    fn encrypt(&self, data: &str) -> String;
    fn decrypt(&self, data: &str) -> String;
}

impl CipherTrait for Cipher {
    fn new(key: String) -> Self {
        let key = hex::decode(key).unwrap();
        let boxed_slice = key.into_boxed_slice();
        let boxed_array: Box<[u8; 32]> = match boxed_slice.try_into() {
            Ok(ba) => ba,
            Err(_) => panic!("Key must be 32 bytes long"),
        };
        Cipher { key: *boxed_array }
    }

    fn encrypt(&self, data: &str) -> String {
        let key = &self.key.into();
        let cipher = Aes256Gcm::new(key);
        let nonce = Aes256Gcm::generate_nonce(&mut OsRng);
        let encrypted = cipher.encrypt(&nonce, data.as_ref()).unwrap();
        let nonce_size = [0, 0, 0, 12];
        let encrypted = [&nonce_size, nonce.as_slice(), encrypted.as_ref()].concat();
        BASE64_STANDARD.encode(encrypted)
    }

    fn decrypt(&self, data: &str) -> String {
        let cipher = Aes256Gcm::new(&self.key.into());
        let decoded = BASE64_STANDARD.decode(data.as_bytes()).unwrap();
        let nonce = Nonce::<Aes256Gcm>::from_slice(&decoded[4..16]);
        let data = &decoded[16..];
        String::from_utf8(cipher.decrypt(nonce, data).unwrap()).unwrap()
    }
}