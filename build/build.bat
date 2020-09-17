::fyne package
SET CGO_ENABLED=1
SET CC=x86_64-w64-mingw32-gcc
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w -X 'configs.buildTime=%date:~0,4%-%date:~5,2%-%date:~8,2% %time:~0,2%:%time:~3,2%:%time:~6,2%' -H windowsgui" -o gormat.exe
upx --best gormat.exe
