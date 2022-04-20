package main

import (
	"fmt"

	"github.com/makifdb/dotfryer/cli"
)

func main() {
	cli.ReadConfigFile()
	cli.InstallPackages()
	f := cli.ReadDotfilesFile()
	fmt.Println(f)
}
