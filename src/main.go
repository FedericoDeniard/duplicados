package main

import (
	"bytes"
	"duplicate-files/src/hashes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("\033[1;33mIniciando búsqueda de duplicados\033[0m\n")
	defer func() {
		fmt.Printf("\033[1;32mTiempo total: %v\033[0m\n", time.Since(start))
	}()
	basePath, _ := os.Getwd()
	filesHashes := hashes.HashFiles(basePath)

	classified := hashes.GroupByHashes(filesHashes)
	var output string
	i := 1
	for _, paths := range classified {
		if len(paths) > 1 {
			output += fmt.Sprintf("\033[1;32m%d.\033[0m %v\n", i, paths)
			output += "\033[1;33m-------------------------------------------------------------------------\033[0m\n"
			i++
		}
	}

	showWithPager(output)
}

func showWithPager(output string) error {
	cmd := exec.Command("less", "-R")
	cmd.Stdin = bytes.NewReader([]byte(output))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
