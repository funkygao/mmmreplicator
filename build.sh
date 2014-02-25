#! /bin/bash -e
#===================
# build with BuildId
#===================
cd $(dirname $0)/daemon/mmmd
ID=$(git rev-parse HEAD | cut -c1-7)
go build -ldflags "-X main.BuildID $ID"

#=========
# show ver
#=========
./mmmd -version
