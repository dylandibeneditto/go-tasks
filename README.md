# Tasks
This app is a simple app which was created as a project to learn go. The project is something I needed to keep track of projects and uses a (very basic) source control system with commits to keep track of actions.

## Installation
In order to run this app you can clone the repository and either use the precompiled file and run it in the terminal or compile/interpret the code using the go interpretter. 
*interpretter*
```sh
go run main.go
```
*compiler*
```sh
go build -o tasks main.go
```
For a faster use of the program, place the compiled executable into the `/usr/local/bin` directory (if you are on Mac) and add this line to your `.zshrc`:
```sh
alias tasks="usr/local/bin/tasks"
```