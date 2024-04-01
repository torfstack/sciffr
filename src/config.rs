use toml;
use hex;
use rand;
use rand::Rng;
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize)]
pub struct Config {
    pub encryption: Encryption,
}

#[derive(Deserialize, Serialize)]
pub struct Encryption {
    pub key: String,
}

fn check_sciffr_dir() {
    match home::home_dir() {
        None => panic!("Could not find home directory"),
        Some(dir) => {
            let sciffr_dir = dir.join(".sciffr");
            if !sciffr_dir.exists() {
                match std::fs::create_dir(&sciffr_dir) {
                    Err(e) => panic!("Could not create .sciffr directory: {}", e),
                    Ok(_) => println!("Created .sciffr directory at {}", sciffr_dir.display()),
                }
            }
        }
    }
}

pub(crate) fn read_in_config() -> Config {
    check_sciffr_dir();
    let config_path = home::home_dir().unwrap().join(".sciffr/config.toml");
    if !config_path.exists() {
        match std::fs::File::create(&config_path) {
            Err(e) => panic!("Could not create config file: {}", e),
            Ok(_) => {
                println!("Created config file at {}", config_path.display());
                let key = random_key();
                println!("Generated random key and write to config file");
                let config = Config { encryption: Encryption { key }};
                let toml = toml::to_string(&config).unwrap();
                std::fs::write(&config_path, toml).unwrap();
            },
        }
    }
    let config = std::fs::read_to_string(&config_path).unwrap();
    toml::from_str(&config).unwrap()
}

fn random_key() -> String {
    let key = rand::thread_rng()
        .sample_iter(&rand::distributions::Standard)
        .take(32)
        .collect::<Vec<u8>>();
    return hex::encode(key);
}