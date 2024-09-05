# Tasks CLI

**Tasks** is a simple command-line interface (CLI) task manager built with Go and Cobra.

## Features

- Add new tasks
- List all tasks
- Mark tasks as complete
- Delete tasks

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Harryalways317/tasks.git
    cd tasks
    ```

2. **Build the project**:

    ```bash
    make build
    ```

## Usage

### Add a New Task

Add a new task with a description:

```bash
./tasks add "task description"
```

### List All Tasks

View all active tasks:

```bash
./tasks list
```

### List All Tasks (Including Completed Ones)

View both active and completed tasks:

```bash
./tasks list --all
```
