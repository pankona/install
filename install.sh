#!/bin/bash -e

mkdir /tmp/install
cd /tmp/install

# This example uses 0.0.4 but let's choose latest version
wget https://github.com/pankona/install/releases/download/v0.0.4/install_0.0.4_Linux_x86_64.tar.gz
tar zxf install_0.0.4_Linux_x86_64.tar.gz

./install asdf
./install bashrc
source ~/.bashrc
./install all
