package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type app struct {
	name, when, why string
}

func main() {
	welvet := os.Getenv("WELVET_ROOT")
	if welvet == "" {
		wd, _ := os.Getwd()
		welvet = filepath.Clean(filepath.Join(wd, "../../../welvet"))
	}
	appsRoot := filepath.Join(welvet, "apps")

	apps := []app{
		{"octo", "chat / agent CLI over transformer entities", "main interactive shell"},
		{"flux2", "image generation experiments", "diffusion-style app shell"},
		{"mosstts", "speech synthesis demos", "TTS product path"},
		{"kokoro", "Kokoro TTS (placeholder — empty dir)", "reserved; use mosstts/qwentts until filled"},
		{"aai", "AAI / Lucy organism demos", "cameral + lucy product"},
		{"biotalk", "bio / organism chat surfaces", "domain UI over welvet"},
		{"qwenasr", "Qwen ASR", "speech recognition app"},
		{"qwentts", "Qwen TTS", "speech synthesis app"},
		{"planetbridging", "planetbridging demos", "bridging experiments"},
	}

	fail := 0
	fmt.Println("Apps never import into the engine — run each from its own go.mod.\n")
	for _, a := range apps {
		dir := filepath.Join(appsRoot, a.name)
		st, err := os.Stat(dir)
		if err != nil || !st.IsDir() {
			fmt.Printf("FAIL %-16s missing under apps/\n", a.name)
			fail++
			continue
		}
		hasMod := "no go.mod"
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			hasMod = "go.mod ok"
		}
		// kokoro is an intentional empty placeholder
		ents, _ := os.ReadDir(dir)
		note := hasMod
		if a.name == "kokoro" && len(ents) == 0 {
			note = "empty placeholder"
		}
		fmt.Printf("OK   %-16s [%s]\n", a.name, note)
		fmt.Printf("     when: %s\n", a.when)
		fmt.Printf("     where: apps/%s\n", a.name)
		fmt.Printf("     why:  %s\n\n", a.why)
	}
	if fail > 0 {
		os.Exit(1)
	}
	fmt.Println("All app folders present. Deep cameral proofs: example/cam/")
}
