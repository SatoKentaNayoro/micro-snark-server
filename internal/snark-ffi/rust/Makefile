SHELL=/usr/bin/env bash

dev: snark_ffi_dev move_lib
release: snark_ffi_release move_lib

.PHONY: move_lib
move_lib:
	@mv target/*/libsnarkcrypto.a ../

.PHONY: snark_ffi_dev
snark_ffi_dev: src/lib.rs Cargo.toml
	cargo +nightly build

.PHONY: snark_ffi_release
snark_ffi_release: src/lib.rs Cargo.toml
	cargo +nightly build --release

clean:
	rm -rf target
	rm ../libsnarkcrypto.a