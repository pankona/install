#!/bin/bash -e

mkdir /tmp/install
cd /tmp/install

wget https://github.com/pankona/install/releases/download/v0.0.7/install_0.0.7_Linux_x86_64.tar.gz
tar zxf install_0.0.7_Linux_x86_64.tar.gz

./install -asdf
./install -bashrc
source ~/.bashrc
./install -all
