package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	welvet := os.Getenv("WELVET_ROOT")
	if welvet == "" {
		wd, _ := os.Getwd()
		welvet = filepath.Clean(filepath.Join(wd, "../../../welvet"))
	}
	modelDir := filepath.Join(welvet, "model", "kokoro")
	appsDir := filepath.Join(welvet, "apps", "kokoro")

	me, err := os.ReadDir(modelDir)
	must(err)
	ae, err := os.ReadDir(appsDir)
	must(err)

	fmt.Println("When: you want Kokoro-style TTS")
	fmt.Println("Where: planned at model/kokoro + apps/kokoro — both empty today")
	fmt.Println("Why: placeholders reserved; working TTS apps are mosstts / qwentts")
	fmt.Printf("model/kokoro entries: %d\n", len(me))
	fmt.Printf("apps/kokoro entries:  %d\n", len(ae))
	if len(me) != 0 || len(ae) != 0 {
		panic("expected empty kokoro placeholders")
	}

	for _, name := range []string{"mosstts", "qwentts"} {
		dir := filepath.Join(welvet, "apps", name)
		st, err := os.Stat(dir)
		if err != nil || !st.IsDir() {
			panic(name + " missing")
		}
		fmt.Printf("OK working TTS shell: apps/%s\n", name)
	}
	fmt.Println("OK — kokoro slots empty; use mosstts/qwentts until kokoro lands")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
