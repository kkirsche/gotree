#!/bin/bash

mkdir -p bin
for os in darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows; do
  for arch in "386" amd64; do
    echo "[+] Building for OS: ${os} Architecture: ${arch}"
    if [ "${os}" == "windows" ]; then
      env GOOS="${os}" GOARCH="${arch}" go build -v -o bin/gotree-${os}-${arch}.exe
    else
      env GOOS="${os}" GOARCH="${arch}" go build -v -o bin/gotree-${os}-${arch}
    fi
  done
done
