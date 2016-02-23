//
// file: ./clean.go
// desc: binary for remove os generated files
// author:
// legal:
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type stringSlice []string

func (i *stringSlice) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *stringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	all   = flag.Bool("all", false, "which include subdirectory")
	patternVar stringSlice
)

func main() {
	flag.Var(&patternVar, "pattern", "which add pattern")

	flag.Parse()

	patterns := []string{
		".DS_Store",
		".DS_Store?",
		"._*",
		".Spotlight-V100",
		".Trashes",
		"ehthumbs.db",
		"Thumbs.db",
	}
	if flag.NFlag() > 0 {
		for i := 0; i < len(patternVar); i++ {
			patterns = append(patterns, patternVar[i])
		}
	}

	if len(flag.Args()) == 0 {
		clean("./", patterns)
	} else {
		for _, args := range flag.Args() {
			clean(args, patterns)
		}
	}

}

func clean(dir string, patterns []string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		hasPattern, pattern := hasPattern(f.Name(), patterns)
		if !f.IsDir() && hasPattern == 0 {
			err := os.Remove(filepath.Clean(dir + "/" + f.Name()))
			if err != nil {
				log.Print(err)
			} else {
				log.Println("Remove", filepath.Clean(dir+"/"+f.Name()))
			}
		} else if !f.IsDir() && hasPattern == 1 {
			continue
		} else if !f.IsDir() && hasPattern == 2 {
			globFiles, err := filepath.Glob(dir+"/"+pattern)
			if err!= nil {
				log.Println(err)
			} else {
				for _, globFile := range globFiles {
					err2 := os.Remove(filepath.Clean(globFile))
					if err2 != nil {
						log.Print(err2)
					} else {
						log.Println("Remove", filepath.Clean(globFile))
					}
				}
			}
		} else if *all {
			clean(filepath.Clean(dir+"/"+f.Name()), patterns)
		}
	}
}

func hasPattern(name string, patterns []string) (int, string) {
	for _, pattern := range patterns {
		if strings.Contains(pattern, "*") {
			return 2, pattern
		} else if strings.Contains(name, pattern) {
			return 0, ""
		}
	}
	return 1, ""
}
