#!/bin/bash 
#

type=$1 

alias znvm="docker run --rm -it -v $PWD:/d2vm -w /d2vm linkacloud/d2vm:latest"

cat <<! |base64 -d >Dockerfile
RlJPTSBidXN5Ym94CgpDT1BZIGludGVyZmFjZXMgL2V0Yy9uZXR3b3JrCgpSVU4gYXBrIGFkZCAt
LW5vLWNhY2hlIGFnZXR0eSBvcGVucmMgXAogICAgJiYgcmMtdXBkYXRlIGFkZCBsb2NhbCBkZWZh
dWx0IFwKICAgICYmIGVjaG8gcmNfbG9nZ2VyPVwiTk9cIiA+PiAvZXRjL3JjLmNvbmYgXAogICAg
JiYgZWNobyByY192ZXJib3NlPVwiTk9cIiA+PiAvZXRjL3JjLmNvbmYgCgpDT1BZIGltZy5zdGFy
dCAvZXRjL2xvY2FsLmQgCgpDT1BZIGluaXR0YWIgL2V0YwoK
!

docker build . -t unikernel
./znvm convert localhost/unikernel -o unikernel "$type"
