#!/usr/bin/env bash

# Install dependent
apt update
apt upgrade

apt install -y curl unzip wget snapd

# Install Docker
curl -sSL https://get.docker.com/ | sh

# Install Golang SDK
snap install go --classic

# Install Dart SDK
wget https://storage.googleapis.com/dart-archive/channels/stable/release/latest/linux_packages/dart_2.7.1-1_amd64.deb
dpkg -i dart_2.7.1-1_amd64.deb
rm dart_2.7.1-1_amd64.deb
echo 'export PATH="$PATH:/usr/lib/dart/bin"' >> ~/.profile
