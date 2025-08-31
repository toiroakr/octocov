package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"

    "github.com/k1LoW/octocov/coverage"
    "github.com/k1LoW/octocov/report"
)

func main() {
    in := flag.String("in", "coverage.out", "path to coverage.out")
    out := flag.String("out", "report.json", "path to report.json")
    repo := flag.String("repo", "local/repo", "repository in owner/repo form")
    flag.Parse()

    cov, _, err := coverage.NewGocover().ParseReport(*in)
    if err != nil {
        log.Fatalf("parse coverage: %v", err)
    }

    r, err := report.New(*repo)
    if err != nil {
        log.Fatalf("create report: %v", err)
    }
    r.Coverage = cov

    if err := os.MkdirAll(filepath.Dir(*out), 0755); err != nil { // #nosec
        log.Fatalf("make dir: %v", err)
    }
    if err := os.WriteFile(*out, r.Bytes(), 0644); err != nil { // #nosec
        log.Fatalf("write report: %v", err)
    }
}
