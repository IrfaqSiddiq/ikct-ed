#!/bin/bash
rm ikct_ed ikct_ed.tar.gz
/usr/local/go/bin/go get
GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o ./ikct_ed
#upx ikct_ed
tar -cvzf ikct_ed.tar.gz makefile ikct_ed .env.sample
