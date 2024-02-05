#! /bin/sh
echo "Running go build command"
GOOS=windows GOARCH=amd64 go build -o bin/interset.exe main.go actions.go

echo "Setting permission for binary file in bin"
chmod +x bin/interset.exe

echo "Running the binary file in bin"
bin/interset.exe A_f.csv B_f.csv