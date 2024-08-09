package flakes

import (
	"embed"
	"os"
	"os/exec"

	"errors"
)

func Update() ([]byte, error) {
	return exec.Command("/usr/bin/env", "nix", "flake", "update").CombinedOutput()
}

func UpdateInput(name string) ([]byte, error) {
	return exec.Command("/usr/bin/env", "nix", "flake", "lock", "--update-input", name).CombinedOutput()
}

//go:embed all:templates
var fs embed.FS

func Init(templ string) error {
	switch templ {
	case "Go":
		templbts, _ := fs.ReadFile("templates/go/flake.nix")
		templs := string(templbts)
		return os.WriteFile("flake.nix", []byte(templs), 0644)
	}
	return errors.New("Idk what is happened but there is no template " + templ)
}
