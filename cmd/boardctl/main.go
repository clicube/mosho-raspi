package main

import (
	"fmt"
	"os"

	"mosho-raspi/presentation"
)

func main() {
	if err := presentation.CmdlineExecute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
