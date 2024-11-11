package main

import (
	"log"

	"github.com/kuwairo/inaneshell/shell"
)

func main() {
	var sh shell.Shell
	if err := sh.Load(shell.ConfigPath); err != nil {
		log.Fatalln(err)
	}
	if err := sh.ChangeDirectory([]string{"cd"}); err != nil {
		log.Fatalln(err)
	}
	if err := sh.Loop(); err != nil {
		log.Fatalln(err)
	}
}
