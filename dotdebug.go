package main

import (
	"debug/elf"
	"fmt"
	"os"
)

const buildIdLen = 20

func argh(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fd, err := os.Open("/bin/bash")
	argh(err)
	bash, err := elf.NewFile(fd)
	argh(err)
	section := bash.Section(".note.gnu.build-id")
	if section == nil {
		panic("No note")
	}
	bytes, err := section.Data()
	argh(err)
	fmt.Printf("%x\n", bytes[len(bytes)-buildIdLen:])
}
