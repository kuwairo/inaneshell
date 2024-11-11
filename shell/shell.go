package shell

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	cfg config
}

func (s *Shell) Prompt() {
	fmt.Printf("\n%s", s.cfg.prompt)
}

func (s *Shell) Loop() error {
	s.Prompt()
	scanner := bufio.NewScanner(os.Stdin)
	for ; scanner.Scan(); s.Prompt() {
		line := strings.TrimSpace(scanner.Text())
		if line == s.cfg.exitCommand {
			break
		}
		cmd := strings.Fields(line)
		if len(cmd) < 1 {
			continue
		}

		if s.cfg.allowCD && cmd[0] == "cd" {
			if err := s.ChangeDirectory(cmd); err != nil {
				fmt.Printf("cd: %s\n", err)
			}
			continue
		}

		path, ok := s.cfg.execs[cmd[0]]
		if !ok {
			fmt.Printf("%s: command not found\n", cmd[0])
			continue
		}
		if err := s.exec(path, cmd); err != nil {
			var exitError *exec.ExitError
			if errors.As(err, &exitError) {
				fmt.Println(exitError)
			} else {
				fmt.Printf("%s: %s\n", path, err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read line: %w", err)
	}

	return nil
}

func (s *Shell) ChangeDirectory(cmd []string) (err error) {
	var dir string
	switch len(cmd) {
	case 1:
		dir, err = os.UserHomeDir()
		if err != nil {
			dir, err = os.TempDir(), nil
		}
	case 2:
		dir = cmd[1]
	default:
		return errors.New("too many arguments")
	}

	if err := os.Chdir(dir); err != nil {
		return errors.New("no such file or directory")
	}

	return nil
}

func (s *Shell) exec(path string, cmd []string) error {
	var args []string
	if len(cmd) > 1 {
		args = cmd[1:]
	}

	c := exec.Command(path, args...)
	c.Stdout, c.Stderr = os.Stdout, os.Stdout

	return c.Run()
}
