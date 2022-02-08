all: target/debug/libsnarkcrypto.a

target/debug/libsnarkcrypto.a: src/lib.rs Cargo.toml
	cargo +nightly build

clean:
	rm -rf target