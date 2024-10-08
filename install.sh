#!/bin/bash -ex

cd "$(dirname "$(readlink -f "$0")")"

rm -rf /tmp/install-${V}
mkdir /tmp/install-${V}
cd /tmp/install-${V}
V="0.0.35"

wget https://github.com/pankona/install/releases/download/v${V}/install_${V}_Linux_x86_64.tar.gz
tar zxf install_${V}_Linux_x86_64.tar.gz

./install -apttools
./install -asdf
./install -bashrc
source ~/mybashrc
# asdf update --head # to support asdf plugin-add in alpine
# ./install -ruby
./install -nodejs
# ./install -yarn
./install -ghq
./install -fzf
./install -git
./install -golang
./install -deno
./install -vim
./install -gh
# ./install -docker
./install -prettier
# ./install -kubectl
