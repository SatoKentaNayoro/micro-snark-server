all: make_rust make_go

make_rust:
	$(MAKE) -C ./snark-ffi/rust

make_go:
	go build

clean:
	$(MAKE) -C ./snark-ffi/rust clean