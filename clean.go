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
	"regexp"
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
	regex stringSlice
)

func main() {
	flag.Var(&regex, "regex", "which add pattern with regex")

	flag.Parse()

	patterns := []string{
		".DS_Store",
		".DS_Store?",
		"^_.*",
		".Spotlight-V100",
		".Trashes",
		"ehthumbs.db",
		"Thumbs.db",
	}
	if flag.NFlag() > 0 {
		for i := 0; i < len(regex); i++ {
			patterns = append(patterns, regex[i])
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
		hasPattern := hasPattern(f.Name(), patterns)
		if !f.IsDir() && hasPattern {
			err := os.Remove(filepath.Clean(dir + "/" + f.Name()))
			if err != nil {
				log.Print(err)
			} else {
				log.Println("Remove", filepath.Clean(dir+"/"+f.Name()))
			}
		} else if !f.IsDir() && !hasPattern {
			continue
		} else if *all {
			clean(filepath.Clean(dir+"/"+f.Name()), patterns)
		}
	}
}

func hasPattern(name string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(pattern, "*") {
			matched, err := regexp.MatchString(pattern, name)
			if err != nil {
				log.Panic(err)
			} else if matched {
				return true
			}
		} else if strings.Contains(name, pattern) {
			return true
		}
	}
	return false
}
