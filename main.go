package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func goDir(subpaths ...string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	return filepath.Join(append([]string{usr.HomeDir, "go"}, subpaths...)...)
}

func getClosestMatches(target string, list []string) []string {
	matches := map[string]bool{}
	pkgName := func(s string) string {
		lastParts := strings.Split(s, "/")
		last := lastParts[len(lastParts)-1]
		return strings.Split(last, "@")[0]
	}
	for _, item := range list {
		if pkg := pkgName(item); target == pkg {
			pkgNoVersion := strings.Split(item, "@")[0]
			matches[pkgNoVersion] = true
		}
	}
	var res []string
	for k := range matches {
		res = append(res, k)
	}
	return res
}

func main() {
	binDir := goDir("bin")
	pkgDir := goDir("pkg", "mod")

	binaries, err := os.ReadDir(binDir)
	if err != nil {
		log.Fatalf("Failed to read bin directory: %v", err)
	}

	var packages []string
	err = filepath.Walk(pkgDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "cache" {
			return filepath.SkipDir
		}
		if info.IsDir() && strings.Contains(info.Name(), "@") {
			installPkgDir := strings.Replace(path, pkgDir, "", 1)
			installPkgDir = strings.TrimLeft(installPkgDir, "/")
			packages = append(packages, installPkgDir)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to read pkg directory: %v", err)
	}

	type match struct {
		binary fs.DirEntry
		pkgs   []string
	}

	var none, one, some []match
	for _, binary := range binaries {
		if binary.IsDir() {
			continue
		}
		matches := getClosestMatches(binary.Name(), packages)
		m := match{binary, matches}
		switch len(matches) {
		case 0:
			none = append(none, m)
		case 1:
			one = append(one, m)
		default:
			some = append(some, m)
		}
	}

	println := func(s ...any) { fmt.Println(s...) }
	printf := func(tmpl string, args ...any) { println(fmt.Sprintf(tmpl, args...)) }

	println()
	printf("%d with NO matches", len(none))
	for _, m := range none {
		printf("  %s", m.binary.Name())
	}

	println()
	printf("%d with ONE match", len(one))
	for _, m := range one {
		printf("  %s -> %s", m.binary.Name(), m.pkgs[0])
	}
	if len(one) > 0 {
		var onePkgs []string
		for _, m := range one {
			onePkgs = append(onePkgs, m.pkgs[0])
		}
		cmd := fmt.Sprintf("go install %s", strings.Join(onePkgs, " "))
		println()
		printf("  INSTALL: %s", cmd)
	}

	println()
	printf("%d with MULTIPLE matches", len(some))
	for _, m := range some {
		printf("  %s -> %v", m.binary.Name(), m.pkgs)
	}
}
