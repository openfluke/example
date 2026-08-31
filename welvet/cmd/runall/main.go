// runall smoke-tests every welvet book example that has a main.go.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type chapter struct {
	Slug    string `json:"slug"`
	Num     string `json:"num"`
	Title   string `json:"title"`
	HasMain bool   `json:"has_main"`
}

func main() {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if filepath.Base(root) == "runall" {
		root = filepath.Clean(filepath.Join(root, "../.."))
	}
	metaPath := filepath.Join(root, "_chapters.json")
	raw, err := os.ReadFile(metaPath)
	if err != nil {
		fmt.Println("missing _chapters.json — regenerate READMEs")
		os.Exit(1)
	}
	var chs []chapter
	if err := json.Unmarshal(raw, &chs); err != nil {
		panic(err)
	}

	// Shared caches (same as env.sh)
	cache := filepath.Join(root, ".cache")
	_ = os.MkdirAll(filepath.Join(cache, "gocache"), 0o755)
	_ = os.MkdirAll(filepath.Join(cache, "gotmp"), 0o755)
	_ = os.MkdirAll(filepath.Join(cache, "tmp"), 0o755)

	fail := 0
	skip := 0
	ok := 0
	for _, ch := range chs {
		if !ch.HasMain {
			skip++
			continue
		}
		dir := filepath.Join(root, ch.Slug)
		fmt.Printf("\n######## %s %s ########\n", ch.Num, ch.Slug)
		cmd := exec.Command("go", "run", ".")
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = append(os.Environ(),
			"GOCACHE="+filepath.Join(cache, "gocache"),
			"GOTMPDIR="+filepath.Join(cache, "gotmp"),
			"TMPDIR="+filepath.Join(cache, "tmp"),
		)
		start := time.Now()
		err := cmd.Run()
		fmt.Printf("— %s (%s)\n", ch.Slug, time.Since(start).Round(time.Millisecond))
		if err != nil {
			fmt.Printf("FAILED %s: %v\n", ch.Slug, err)
			fail++
			continue
		}
		ok++
	}
	fmt.Printf("\n======== welvet book: ok=%d fail=%d skip=%d ========\n", ok, fail, skip)
	if fail > 0 {
		os.Exit(1)
	}
}
