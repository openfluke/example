package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Apps import the engine — the engine never imports apps.")
	fmt.Println("Full when/where/why inventory: go run ../75-apps-map")
	wd, _ := os.Getwd()
	apps := filepath.Clean(filepath.Join(wd, "../../../welvet/apps"))
	entries, err := os.ReadDir(apps)
	if err != nil {
		panic(err)
	}
	fmt.Printf("apps root: %s (%d entries)\n", apps, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			fmt.Println(" -", e.Name())
		}
	}
}
