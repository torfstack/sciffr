use aes;
use aes::cipher::KeyInit;
use aes_gcm::aead::{Aead, OsRng};
use aes_gcm::{AeadCore, Aes256Gcm};
use base64::prelude::*;

pub(crate) struct Cipher {
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
        BASE64_STANDARD.encode(encrypted)
    }

    fn decrypt(&self, data: &str) -> String {
        String::from("test")
    }
}