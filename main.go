// SPDX-License-Identifier: 0BSD
// SPDX-FileCopyrightText: 2025 Hajime Hoshi

package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/overlaytest/internal/foo"
)

func main() {
	if err := xmain(); err != nil {
		panic(err)
	}
}

func xmain() error {
	foo.Foo()

	f, err := os.Create("overlay.json")
	if err != nil {
		return err
	}
	defer f.Close()

	barFile, err := os.CreateTemp("", "bar.go")
	if err != nil {
		return err
	}
	defer barFile.Close()

	if _, err := barFile.WriteString(`package foo

func init() {
	message = "Replaced"
}
`); err != nil {
		return err
	}
	if err := barFile.Sync(); err != nil {
		return err
	}

	pkg, err := goPkgDir("github.com/hajimehoshi/overlaytest/internal/foo")
	if err != nil {
		return err
	}
	replaces := map[string]map[string]string{
		"Replace": {
			filepath.Join(pkg, "bar.go"): barFile.Name(),
		},
	}
	if err := json.NewEncoder(f).Encode(replaces); err != nil {
		return err
	}

	return nil
}

func goPkgDir(pkg string) (string, error) {
	cmd := exec.Command("go", "list", "-f", "{{.Dir}}", pkg)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
