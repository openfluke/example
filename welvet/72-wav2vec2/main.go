package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/openfluke/welvet/model/wav2vec2"
)

func main() {
	cfg := wav2vec2.Base960h()
	fmt.Printf("Base960h hidden=%d layers=%d heads=%d vocab=%d\n",
		cfg.HiddenSize, cfg.NumHiddenLayers, cfg.NumAttentionHeads, cfg.VocabSize)

	vocab, err := wav2vec2.LoadVocabBytes([]byte(`{"<pad>":0,"|":1,"E":2,"T":3,"A":4}`), 0)
	must(err)
	fmt.Println("DecodeCTCGreedy:", vocab.DecodeCTCGreedy([]int{4, 4, 3, 2}))

	missing := filepath.Join(os.TempDir(), "welvet-no-wav2vec2-weights")
	_, err = wav2vec2.LoadHFDir(missing)
	if err == nil {
		panic("LoadHFDir should fail without snapshot")
	}
	fmt.Println("LoadHFDir without weights:", err)

	if dir := os.Getenv("WAV2VEC2_DIR"); dir != "" {
		m, err := wav2vec2.LoadHFDir(dir)
		must(err)
		wav := os.Getenv("WAV2VEC2_WAV")
		if wav == "" {
			panic("missing WAV2VEC2_WAV")
		}
		text, err := m.TranscribeFile(wav)
		must(err)
		fmt.Println("transcribe:", text)
	} else {
		fmt.Println("set WAV2VEC2_DIR (+ WAV2VEC2_WAV) for full ASR smoke")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
