# Getting Started with running
This project was created to use enable you run manually or to use a shell scripting

#### Run it manually
While in the root directory run the command:
#### `go run main.go actions.go <file1.csv> <file2.csv>`
Replace the file1.csv and file2.csv with your own two csv files for the two datasets needed

#### Run it using shell scripting (Windows)
On a windows system using a bash shell environment run (I used a git bash in my windows not pure linux):
#### `./run.sh`

#### Run it using shell scripting (Unix Kernel system)
On linux or a unix kernel system using a bash shell environment run:
#### `./run_unix.sh`

#### Testing Full coverage
To run the test and report code coverage information, use the command below in your console/terminal:
#### `go test -cover`

#### Run Go test for specific test function
To run test for a specific test function use the code below in in your console/terminal:
#### `go test -v -run=<test_name>` 
#####  Example: `go test -v -run=TestReadCSVKeys` 

Either ways you should see your output.
Thanks

Enjoy:
