package main

import (
	"bytes"
	customFlags "duplicate-files/src/flags"
	"duplicate-files/src/hashes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	excludedDirsFlag := flag.String("exclude-dirs", "", "Comma-separated list of directories to exclude from search")
	helpFlag := flag.Bool("help", false, "Show help message")
	showHiddenFiles := flag.Bool("show-hidden", false, "Include hidden files and directories in search")
	includeExtensions := flag.String("include-ext", "", "Comma-separated list of file extensions to include (e.g., .jpg,.png)")
	excludeExtensions := flag.String("exclude-ext", "", "Comma-separated list of file extensions to exclude")
	useSHA256Flag := flag.Bool("use-sha256", false, "Use SHA256 for hashing (slower but more secure than default MD5)")

	flag.Parse()
	if *helpFlag {
		flag.Usage()
		return
	}

	customFlags := customFlags.CustomFlags{
		ShowHiddenFiles:        *showHiddenFiles,
		ExcludedRoutes:         strings.Split(*excludedDirsFlag, ","),
		FileExtensions:         strings.Split(*includeExtensions, ","),
		ExcludedFileExtensions: strings.Split(*excludeExtensions, ","),
		UseSHA256:              *useSHA256Flag,
	}
	customFlags.Normalize()

	fmt.Printf("\033[1;33mIniciando búsqueda de duplicados\033[0m\n")

	basePath, _ := os.Getwd()
	filesHashes := hashes.HashFiles(basePath, customFlags)

	var output string
	i := 1
	for _, paths := range filesHashes {
		if len(paths) > 1 {
			output += fmt.Sprintf("\033[1;32m%d.\033[0m %v\n", i, paths)
			output += "\033[1;33m-------------------------------------------------------------------------\033[0m\n"
			i++
		}
	}

	fmt.Printf("\033[1;32mTiempo total: %v\033[0m\n", time.Since(start))
	showWithPager(output)
}

func showWithPager(output string) error {
	cmd := exec.Command("less", "-R")
	cmd.Stdin = bytes.NewReader([]byte(output))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
