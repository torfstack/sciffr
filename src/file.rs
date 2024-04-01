use serde_json;
use serde_json::{Map, Value};
use crate::cipher::{Cipher, CipherTrait};

pub struct FileEncryptor {
    path: String,
    content: Map<String, Value>,
    cipher: Box<dyn CipherTrait>,
}

trait FileEncryptorTrait {
    fn new(file: String, key: String) -> Self;
    fn encrypt(&self, key: String);
}

impl FileEncryptor {
    pub fn new(path: String, key: String) -> Self {
        let file = std::fs::read_to_string(&path).unwrap();
        let content: Map<String, Value> = serde_json::from_str(&file).unwrap();
        let cipher = Box::new(Cipher::new(key));
        FileEncryptor { path, content, cipher}
    }

    pub fn encrypt(&mut self, key: String) {
        let data = self.content.get_mut(&key).unwrap();
        let encrypted = self.cipher.encrypt(data.as_str().unwrap());
        self.content[&key] = Value::from(encrypted);
        let new_content = serde_json::to_string(&self.content).unwrap();
        self.write(new_content);
    }

    fn write(&self, content: String) {
        std::fs::write(&self.path, content).unwrap();
    }
}

