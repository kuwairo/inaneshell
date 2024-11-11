package shell

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigPath = "/etc/inaneshell/config.json"
)

type rawConfig struct {
	Cd      bool     `json:"cd"`
	Exit    string   `json:"exit"`
	Prompt  string   `json:"prompt"`
	Provide []string `json:"provide"`
}

type config struct {
	allowCD     bool
	exitCommand string
	prompt      string
	execs       map[string]string
}

func (s *Shell) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open config: %w", err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var raw rawConfig
	if err := dec.Decode(&raw); err != nil {
		return fmt.Errorf("decode config: %w", err)
	}

	cfg := config{
		allowCD:     raw.Cd,
		exitCommand: raw.Exit,
		prompt:      raw.Prompt,
		execs:       make(map[string]string, len(raw.Provide)),
	}
	for _, path := range raw.Provide {
		cfg.execs[filepath.Base(path)] = path
	}

	// TODO: implement extended validation
	s.cfg = cfg
	return nil
}
