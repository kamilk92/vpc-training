#!/bin/bash

GO_PKG=go1.18.3.linux-amd64.tar.gz
wget -c https://go.dev/dl/$GO_PKG
sudo tar -C /usr/local -xvzf $GO_PKG
rm -rf $GO_PKG
GO_PROJECTS_DIR=go-projects
mkdir -p ~/$GO_PROJECTS_DIR/{bin,src,pkg}
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
echo 'export GOPATH="$HOME/go-projects"' >> ~/.bash_profile
echo 'export GOBIN="$GOPATH/bin"' >> ~/.bash_profile