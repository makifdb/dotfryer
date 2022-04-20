package cli

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func WriteDotfilesFile() {

}

func ReadDotfilesFile() []string {
	dotfiles := []string{}
	f, err := ioutil.ReadFile(C.Location + "dotfiles.list")
	if err != nil {
		fmt.Printf("Error reading dotfiles file: %s\n", err)
		return dotfiles
	}

	for _, line := range strings.Split(string(f), "\n") {
		if line != "" {
			dotfiles = append(dotfiles, line)
		}
	}
	return dotfiles
}
