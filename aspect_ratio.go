package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func getAspectRatio(filePath string) (string, error) {

	type stream struct {
		Streams []struct {
			DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
		} `json:"streams"`
	}

	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	var s stream
	if err := json.NewDecoder(&out).Decode(&s); err != nil {
		return "", err
	}

	if len(s.Streams) == 0 {
		return "", nil
	}

	return s.Streams[0].DisplayAspectRatio, nil
}
