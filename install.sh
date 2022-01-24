#!/bin/bash -e

V="0.0.12"

mkdir /tmp/install
cd /tmp/install

wget https://github.com/pankona/install/releases/download/v${V}/install_${V}_Linux_x86_64.tar.gz
tar zxf install_${V}_Linux_x86_64.tar.gz

./install -asdf
./install -bashrc
source ~/.bashrc
asdf update --head # to support asdf plugin-add in alpine
# ./install -ruby
./install -nodejs
./install -yarn
./install -ghq
./install -git
./install -golang
# ./install -vim
# ./install -gh
# ./install -docker
./install -prettier
# ./install -kubectl
