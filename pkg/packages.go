package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/makifdb/packer"
)

func UpdatePackages() {
	err := packer.Update()
	if err != nil {
		log.Printf("Error updating packages: %s\n", err)
	}
}

func InstallPackages() {
	for _, p := range ReadPackages() {
		err := packer.Install(p)
		if err != nil {
			log.Printf("Error installing package %s: %s\n", p, err)
		}
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
