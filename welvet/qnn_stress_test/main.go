// qnn_stress_test bootstrap — no cgo. Sets PATH + WELVET_QNN_INPROC, then runs ./stress.
//
//	cd example/welvet/qnn_stress_test
//	go run .
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fatal(err)
	}
	dllDir, err := findQNNDllDir(root)
	if err != nil {
		fmt.Println(err)
		fmt.Println("expected: <git>/welvet/native/qnn/windows_arm64/libwelvet_qnn.dll")
		os.Exit(1)
	}

	env := withEnv(os.Environ(), map[string]string{
		"PATH":              dllDir + string(os.PathListSeparator) + os.Getenv("PATH"),
		"WELVET_QNN_INPROC": "1", // skip welvet_qnn_probe.exe child gate
		"WELVET_QNN_DLL_DIR": dllDir,
	})

	cmd := exec.Command("go", "run", "./stress")
	cmd.Dir = root
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			os.Exit(ee.ExitCode())
		}
		fatal(err)
	}
}

func findQNNDllDir(from string) (string, error) {
	candidates := []string{
		filepath.Join(from, "..", "..", "..", "welvet", "native", "qnn", "windows_arm64"),
		filepath.Join(from, "..", "..", "welvet", "native", "qnn", "windows_arm64"),
		filepath.Join(from, "..", "welvet", "native", "qnn", "windows_arm64"),
	}
	if v := os.Getenv("WELVET_QNN_DLL_DIR"); v != "" {
		candidates = append([]string{v}, candidates...)
	}
	for _, c := range candidates {
		abs, err := filepath.Abs(c)
		if err != nil {
			continue
		}
		if _, err := os.Stat(filepath.Join(abs, "libwelvet_qnn.dll")); err == nil {
			return abs, nil
		}
	}
	return "", fmt.Errorf("libwelvet_qnn.dll not found (looked under %s/../../../welvet/native/qnn/windows_arm64)", from)
}

func withEnv(base []string, set map[string]string) []string {
	upper := make(map[string]string, len(set))
	for k, v := range set {
		upper[strings.ToUpper(k)] = v
	}
	out := make([]string, 0, len(base)+len(set))
	seen := map[string]bool{}
	for _, e := range base {
		k, _, ok := strings.Cut(e, "=")
		if !ok {
			out = append(out, e)
			continue
		}
		ku := strings.ToUpper(k)
		if v, hit := upper[ku]; hit {
			if !seen[ku] {
				out = append(out, k+"="+v)
				seen[ku] = true
			}
			continue
		}
		out = append(out, e)
	}
	for k, v := range set {
		ku := strings.ToUpper(k)
		if !seen[ku] {
			out = append(out, k+"="+v)
			seen[ku] = true
		}
	}
	return out
}

func fatal(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
