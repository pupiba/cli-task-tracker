# Task Tracker (cli)

My solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) from [roadmap.sh](https://roadmap.sh/).

## Run project
Clone the repository and navigate to the main.go file:

```bash
git clone https://github.com/pupiba/cli-task-tracker.git
cd cmd/cli-task-tracker
```

Run the following command to build the project:

```bash
go build -o task-cli.exe
```

### Usage example:

Adding a new task:
```bash
./task-cli.exe add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

Updating and deleting tasks:

```bash
./task-cli.exe update 1 "Buy groceries and cook dinner"
./task-cli.exe delete 1
```

Marking a task as in progress or done:
```bash
./task-cli.exe mark-in-progress 1
./task-cli.exe mark-done 1
```

Listing all tasks:

```bash
./task-cli.exe list
```

Listing tasks by status:
```bash
./task-cli.exe list done
./task-cli.exe list todo
./task-cli.exe list in-progress
```
