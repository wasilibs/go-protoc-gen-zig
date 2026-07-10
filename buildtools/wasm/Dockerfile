FROM alpine:3
ARG TARGETARCH

RUN apk add binaryen curl xz
RUN mkdir -p /zig && curl -L https://ziglang.org/download/0.16.0/zig-$(echo $TARGETARCH | sed "s/amd64/x86_64/" | sed "s/arm64/aarch64/")-linux-0.16.0.tar.xz | tar -xJ --strip-components 1 -C /zig

ADD buildtools/wasm/version.txt version.txt
RUN mkdir -p /workspace && curl -L https://github.com/arwalk/zig-protobuf/archive/$(cat version.txt).tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace

# Zig 0.16's smp_allocator trips wasm32-wasi's single-threaded allocator path.
# The generator is a protoc plugin, so stdin/stdout must be streaming under WASI.
RUN sed -i 's/std.heap.smp_allocator/std.heap.wasm_allocator/g; s/File.stdin().reader(/File.stdin().readerStreaming(/g; s/File.stdout().writer(/File.stdout().writerStreaming(/g' bootstrapped-generator/main.zig \
  && /zig/zig build -Dtarget=wasm32-wasi --release=safe

CMD ["cp", "zig-out/bin/protoc-gen-zig.wasm", "/out/"]
