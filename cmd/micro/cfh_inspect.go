package main

import (
    "crypto/sha256"
    "fmt"
    "io"
    "os"
)

// cfhInspect hashes file content without running or interpreting it.
func cfhInspect(path string) (string, int64, error) {
    f, err := os.Open(path)
    if err != nil { return "", 0, err }
    defer f.Close()
    info, err := f.Stat()
    if err != nil { return "", 0, err }
    if !info.Mode().IsRegular() { return "", 0, fmt.Errorf("not a regular file") }
    h := sha256.New()
    size, err := io.Copy(h, f)
    if err != nil { return "", 0, err }
    return fmt.Sprintf("%x", h.Sum(nil)), size, nil
}