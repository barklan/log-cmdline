package pkg

import (
	"fmt"
	"log"

	"github.com/hpcloud/tail"
)

const (
	CMDLINE_PATH = "/media/cmdline"
)

func StartRelaying() {
	log.Println("opening tail session")
	t, err := tail.TailFile(CMDLINE_PATH, tail.Config{Follow: true})
	if err != nil {
		log.Panicln("failed to open tail session on file")
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
