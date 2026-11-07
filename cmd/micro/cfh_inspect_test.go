package main

import (
    "os"
    "path/filepath"
    "testing"
)

func TestCFHInspect(t *testing.T) {
    dir := t.TempDir()
    file := filepath.Join(dir, "sample.txt")
    if err := os.WriteFile(file, []byte("abc"), 0600); err != nil { t.Fatal(err) }
    hash, size, err := cfhInspect(file)
    if err != nil || size != 3 || hash != "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" {
        t.Fatalf("unexpected result: hash=%s size=%d error=%v", hash, size, err)
    }
    if _, _, err := cfhInspect(dir); err == nil { t.Fatal("directory accepted") }
    if _, _, err := cfhInspect(filepath.Join(dir, "missing")); err == nil { t.Fatal("missing file accepted") }
}