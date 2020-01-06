SET CGO_ENABLED=1
SET CC=x86_64-w64-mingw32-gcc
SET GOOS=windows
SET GOARCH=amd64
go build -o main_x64.exe -ldflags="-H windowsgui"
