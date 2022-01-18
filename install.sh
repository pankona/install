#!/bin/bash -e

mkdir /tmp/install
cd /tmp/install

wget https://github.com/pankona/install/releases/download/v0.0.8/install_0.0.8_Linux_x86_64.tar.gz
tar zxf install_0.0.8_Linux_x86_64.tar.gz

./install -asdf
./install -bashrc
source ~/.bashrc

./install -git
./install -golang
./install -vim
./install -gh
./install -prettier

