package main

import (
	_ "embed"

	"github.com/esportzvio/frietorchain/command/root"
	"github.com/esportzvio/frietorchain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
