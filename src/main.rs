use std::io;
use sha2;
use sha2::Digest;
use hex_literal::hex;

fn main() -> io::Result<()> {
    let mut stdin = io::stdin();
    let mut stdout = io::stdout();
    let mut stderr = io::stderr();

    let mut hasher = sha2::Sha256::new();

    io::copy(&mut stdin, &mut hasher)?;

    Ok(())
}
