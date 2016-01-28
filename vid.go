package main

import (
	"fmt"
	"path/filepath"
	"time"
)

type ytVidInfo struct {
	id        string
	published time.Time
	title     string
	desc      string
}

func (vi *ytVidInfo) episodePath() string {
	return filepath.Join(dataSubdirEpisodes,
		fmt.Sprint(vi.id, ".", downloadAudioFormat))
}