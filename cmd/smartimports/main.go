package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

var verbose bool

func main() {
	var targetPath string
	var localPackage string
	var excludedPaths string

	flag.StringVar(&targetPath, "path", ".", "target path to apply smart goimports, can be a file or a directory")
	flag.StringVar(&localPackage, "local", "", "put imports beginning with this string after 3rd-party packages; comma-separated list")
	flag.StringVar(&excludedPaths, "exclude", "", "paths which should be excluded from processing; comma-separated list")
	flag.BoolVar(&verbose, "v", false, "verbose output")

	flag.Parse()

	opts := &imports.Options{
		TabIndent:  true,
		FormatOnly: true,
	}
	imports.LocalPrefix = localPackage

	excludedPathsList := strings.Split(excludedPaths, ",")
	filteredExcludedPaths := make([]string, 0, len(excludedPathsList))
	for _, path := range excludedPathsList {
		path = strings.TrimSpace(path)
		if path == "" {
			continue
		}
		filteredExcludedPaths = append(filteredExcludedPaths, path)
	}

	processDir(targetPath, opts, filteredExcludedPaths)
}

func processDir(path string, opts *imports.Options, excludedPaths []string) error {
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if verbose {
			fmt.Println("processing path", path)
		}
		for _, excludedPath := range excludedPaths {
			if strings.HasPrefix(path, excludedPath) {
				fmt.Println("   skipped because matched this excluded path:", excludedPath)
				return nil
			}
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		return processFile(path, info, opts)
	})
}

func processFile(filename string, info fs.FileInfo, opts *imports.Options) error {
	rawData, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "os.ReadFile")
	}

	res, err := processData(rawData, opts)
	if err != nil {
		return errors.Wrap(err, "processData")
	}

	err = os.WriteFile(filename, res, info.Mode())
	if err != nil {
		return errors.Wrap(err, "os.WriteFile")
	}
	return nil
}

func processData(src []byte, opts *imports.Options) ([]byte, error) {
	res, err := imports.Process("", src, opts)
	if err != nil {
		return nil, errors.Wrap(err, "imports.Process 1")
	}

	res = removeImportEmptyLines(res)

	res, err = imports.Process("", res, opts)
	if err != nil {
		return nil, errors.Wrap(err, "imports.Process 2")
	}

	return res, nil
}

func removeImportEmptyLines(src []byte) []byte {
	r := bufio.NewReader(bytes.NewBuffer(src))
	w := bytes.NewBuffer(make([]byte, 0, len(src)))

	importsStarted := false
	importsEnded := false

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if importsStarted {
			if !importsEnded {
				if strings.TrimSpace(line) == "" {
					continue
				}
				if strings.HasPrefix(line, ")") {
					importsEnded = true
				}
			}
		} else {
			if strings.HasPrefix(line, "import (") {
				importsStarted = true
			}
		}

		_, _ = w.WriteString(line)
	}

	return w.Bytes()
}
