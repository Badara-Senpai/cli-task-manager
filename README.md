# CLI Task Manager 📋

A command-line task management application built in Go using Cobra CLI framework and BoltDB for data persistence. 

## Features

- Easy-to-use command-line interface
- Persistent storage using BoltDB
- Basic task management operations:
    - Add new tasks
    - List all tasks
    - Mark tasks as completed
    - Delete tasks

## Technologies

- Go 1.21.6
- [Cobra](https://github.com/spf13/cobra) - Modern CLI framework
- [BoltDB](https://github.com/boltdb/bolt) - Embedded key/value database
- [go-homedir](https://github.com/mitchellh/go-homedir) - Cross-platform user home directory detection

## Installation

```bash
# Clone the repository
git clone https://github.com/Badara-Senpai/cli_task_manager.git

# Navigate to project directory
cd cli_task_manager

# Install the binary
go install ./

# Make sure your Go bin directory is in your PATH
export PATH=$PATH:$GOPATH/bin
````


## Usage
```
$ task
Task is a CLI task manager

Usage:
  task [command]

Available Commands:
  add         Adds a task to your task list
  completion  Generate the autocompletion script for the specified shell
  do          Marks a task as completed
  help        Help about any command
  list        List all of your tasks.

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.

$ task list
You have no task! Holidays, we are coming 🏖️

$ task add "Learn to solve Rubix Cube"
Added "Learn to solve Rubix Cube" to your task list.

$ task add Finish the Readme file
Added "Finish the Readme file" to your task list.

$ task list
You have the following tasks:
1. Learn to solve Rubix Cube
2. finish the Readme file

$ task do 2
"finish the Readme file" marked as completed.
```

## Project Structure
```bash
cli_task_manager/
├── cmd/
│   ├── root.go    # Root command definition
│   ├── add.go     # Add command implementation
│   └── list.go    # List command implementation
│   └── do.go      # Mark as completed command implementation
├── db/
│   └── db.go      # Database operations
├── main.go        # Application entry point
└── go.mod         # Go module definition
```

