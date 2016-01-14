//
// file: ./clean.go
// desc: binary for remove os generated files
// author:
// legal:
//
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	all     = flag.Bool("all", false, "which include subdirectory")
	pattern = flag.String("pattern", "", "which include subdirectory")
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		clean("./")
	} else {
		for _, args := range flag.Args()[1:] {
			clean(args)
		}
	}

}

func clean(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if !f.IsDir() && hasPattern(f.Name()) {
			err := os.Remove(filepath.Clean(dir + "/" + f.Name()))
			if err != nil {
				log.Print(err)
			} else {
				log.Println("Remove", filepath.Clean(dir+"/"+f.Name()))
			}
		} else if !f.IsDir() && !hasPattern(f.Name()) {
			continue
		} else {
			if *all {
				clean(filepath.Clean(dir + "/" + f.Name()))
			}
		}
	}
}

func hasPattern(name string) bool {
	patterns := []string{
		".DS_Store",
		".DS_Store?",
		"._*",
		".Spotlight-V100",
		".Trashes",
		"ehthumbs.db",
		"Thumbs.db",
	}

	for _, pattern := range patterns {
		if strings.Contains(name, pattern) {
			return true
		}
	}
	return false
}
