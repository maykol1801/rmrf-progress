package main

import (
    "bufio"
    "flag"
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
    "strings"
    "sync/atomic"

    "github.com/vbauerster/mpb/v8"
    "github.com/vbauerster/mpb/v8/decor"
)

func main() {
    // CLI flags
    dryRun := flag.Bool("dry-run", false, "Show what would be deleted without deleting")
    verbose := flag.Bool("verbose", false, "Print each path being deleted or would be deleted")
    force := flag.Bool("force", false, "Skip confirmation prompt for real deletion")
    flag.Parse()

    if flag.NArg() < 1 {
        fmt.Println("Usage: rmrf-progress [--dry-run] [--verbose] [--force] <target-dir>")
        os.Exit(1)
    }
    target := flag.Arg(0)

    // Safety check: prevent deleting / or .
    absTarget, _ := filepath.Abs(target)
    if absTarget == "/" || absTarget == "." {
        fmt.Println("Refusing to delete root or current directory!")
        os.Exit(1)
    }

    // Confirmation prompt for real deletion (skipped if --force or dry-run)
    if !*dryRun && !*force {
        fmt.Printf("Are you sure you want to delete '%s' and all its contents? [y/N]: ", target)
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input != "y" && input != "yes" {
            fmt.Println("Aborted")
            os.Exit(0)
        }
    }

    // Collect all paths in a slice
    var paths []string
    err := filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            fmt.Printf("Warning: cannot access %s: %v\n", path, err)
            return nil
        }
        paths = append(paths, path)
        return nil
    })
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
        os.Exit(1)
    }

    total := int64(len(paths))
    if total == 0 {
        fmt.Println("Nothing to delete")
        return
    }

    // Setup progress bar with ETA decorator
    p := mpb.New()
    var deleted int64
    bar := p.AddBar(total,
        mpb.PrependDecorators(
            decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
        ),
        mpb.AppendDecorators(
            decor.Percentage(decor.WCSyncWidth),
            decor.EwmaETA(decor.ET_STYLE_GO, 30),
        ),
    )

    // Delete items in reverse order (files first, then directories)
    for i := len(paths) - 1; i >= 0; i-- {
        path := paths[i]
        if *dryRun {
            if *verbose {
                fmt.Printf("Would remove: %s\n", path)
            }
        } else {
            err := os.Remove(path)
            if err != nil {
                fmt.Printf("\nFailed to remove %s: %v\n", path, err)
            } else if *verbose {
                fmt.Printf("Removed: %s\n", path)
            }
        }
        atomic.AddInt64(&deleted, 1)
        bar.IncrBy(1)
    }

    p.Wait()

    if *dryRun {
        fmt.Printf("Dry run complete. No files were deleted. %d items would have been removed.\n", deleted)
    } else {
        fmt.Printf("Done! Successfully deleted %d items.\n", deleted)
    }
}
