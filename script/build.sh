CURRENT=$(cd $(dirname $0);pwd)
echo $CURRENT
cd ../

# win用にbuild
GOOS=windows GOARCH=amd64 go build -o dist/ags.exe ./src/main.go
# mac用にbuild
go build -o dist/ags ./src/main.go