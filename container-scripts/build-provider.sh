#!/bin/bash

cd /root/go/src/github.com/billtrust/meetup-terraform-provider/terraform-provider-btmascot
pwd
go get
rm -rf  bin/*
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/terraform-provider-btmascot
