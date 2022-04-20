package cli

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/makifdb/packer"
)

func UpdatePackages() {
	packer.Update()
}

func InstallPackages() {
	for _, p := range ReadPackages() {
		packer.Install(p)
	}

}

func ReadPackages() []string {
	packages := []string{}
	f, err := ioutil.ReadFile(C.Location + "packages.list")
	if err != nil {
		fmt.Printf("Error reading packages file: %s\n", err)
		return packages
	}

	for _, line := range strings.Split(string(f), "\n") {
		if line != "" {
			packages = append(packages, line)
		}
	}
	return packages
}
