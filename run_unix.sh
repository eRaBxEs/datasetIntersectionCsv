#!/bin/bash

echo "Running go build command"
GOOS=linux GOARCH=amd64 go build -o bin/interset main.go actions.go

echo "Setting permission for binary file in bin"
chmod +x bin/interset

echo "Running the binary file in bin"
bin/interset A_f.csv B_f.csv
