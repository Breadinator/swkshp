@echo off
:: For cross compile, cgo doesn't support
SET CGO_ENABLED=0
:: 32-bit shouldn't work because of the size of the workshop IDs
SET GOARCH=amd64

SET GOOS=linux
go build -o build/swkshp-linux .

SET GOOS=darwin
go build -o build/swkshp-darwin .

SET GOOS=windows
go build -o build/swkshp.exe .