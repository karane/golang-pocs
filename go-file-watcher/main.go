package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/fsnotify/fsnotify"
)

func main() {
    // Define a command-line flag for the directory to watch
    dirToWatch := flag.String("dir", "", "Directory to watch for file system events")
    flag.Parse()

    if *dirToWatch == "" {
        log.Fatal("You must specify a directory to watch using the -dir flag")
    }

    // Check if directory exists, create if not
    if _, err := os.Stat(*dirToWatch); os.IsNotExist(err) {
        err := os.MkdirAll(*dirToWatch, 0755)
        if err != nil {
            log.Fatalf("Failed to create directory %s: %v", *dirToWatch, err)
        }
        fmt.Printf("Directory %s created.\n", *dirToWatch)
    }

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)

    // Catch system interrupt signals for graceful shutdown
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                go handleEvent(event)
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            case <-sigs:
                fmt.Println("\nReceived interrupt signal, shutting down...")
                done <- true
                return
            }
        }
    }()

    fmt.Printf("Watching directory: %s\n", *dirToWatch)

    err = watcher.Add(*dirToWatch)
    if err != nil {
        log.Fatal(err)
    }

    <-done
    fmt.Println("Watcher stopped.")
}

func handleEvent(event fsnotify.Event) {
    fmt.Printf("Event: %s\n", event)

    if event.Op&fsnotify.Create == fsnotify.Create {
        fmt.Printf("File created: %s\n", event.Name)
    }
    if event.Op&fsnotify.Write == fsnotify.Write {
        fmt.Printf("File modified: %s\n", event.Name)
    }
    if event.Op&fsnotify.Remove == fsnotify.Remove {
        fmt.Printf("File removed: %s\n", event.Name)
    }
    if event.Op&fsnotify.Rename == fsnotify.Rename {
        fmt.Printf("File renamed: %s\n", event.Name)
    }
    if event.Op&fsnotify.Chmod == fsnotify.Chmod {
        fmt.Printf("File chmod: %s\n", event.Name)
    }
}
