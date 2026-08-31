// runall smoke-tests every cam example package.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// allow running from cam/ or cam/cmd/runall
	if filepath.Base(root) == "runall" {
		root = filepath.Clean(filepath.Join(root, "../.."))
	}
	pkgs := []string{
		"./01_modes", "./02_combine", "./03_camsync", "./04_kit",
		"./05_layers", "./06_recipes",
	}
	fail := 0
	for _, p := range pkgs {
		fmt.Printf("\n######## %s ########\n", p)
		cmd := exec.Command("go", "run", p)
		cmd.Dir = root
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("FAILED %s: %v\n", p, err)
			fail++
		}
	}
	if fail > 0 {
		fmt.Printf("\n%d package(s) failed\n", fail)
		os.Exit(1)
	}
	fmt.Println("\n======== all cam examples ok ========")
}
