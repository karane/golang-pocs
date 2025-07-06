
# Go File Watcher

A simple Go program to watch filesystem events on a specified directory using [fsnotify/fsnotify](https://github.com/fsnotify/fsnotify).

## Features

- Watches file system events like create, write, rename, remove, chmod.
- Handles events concurrently using goroutines.
- Graceful shutdown on interrupt signals.
- Directory to watch is configurable via command-line flag.

## Requirements

- Go 1.16 or later

## Setup

Initialize Go modules and get dependencies:

```bash
go mod tidy
```

## Usage

Run the program specifying the directory to watch with the `-dir` flag:

```bash
go run main.go -dir=./watched-folder
```

If the directory does not exist, it will be created automatically.

## Testing the Watcher

Perform the following operations inside the watched directory to see the watcher detect and print events:

1. **Create a new file:**

```bash
echo "Hello World" > watched-folder/testfile.txt
```

2. **Modify the file:**

```bash
echo "Additional line" >> watched-folder/testfile.txt
```

3. **Rename the file:**

```bash
mv watched-folder/testfile.txt watched-folder/testfile_renamed.txt
```

4. **Change file permissions:**

```bash
chmod 644 watched-folder/testfile_renamed.txt
```

5. **Remove the file:**

```bash
rm watched-folder/testfile_renamed.txt
```

6. **Create a subdirectory:**

```bash
mkdir watched-folder/subdir
```

7. **Stop the watcher**

Press `Ctrl+C` in the terminal to stop the program gracefully.
