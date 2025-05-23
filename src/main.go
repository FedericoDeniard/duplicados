package main

import (
	"bytes"
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
	var excludedRoutes []string
	excludedRoutesFlag := flag.String("exclude", "", "Rutas a excluir separadas por comas")
	var helpFlag = flag.Bool("help", false, "Mostrar ayuda")
	var showHiddenFiles = flag.Bool("show-hidden", false, "Mostrar archivos ocultos")
	var fileExtensions []string
	fileExtensionsFlag := flag.String("file-extensions", "", "Extensiones a buscar separadas por comas")
	var excludedFileExtensions []string
	excludedFileExtensionsFlag := flag.String("exclude-file-extensions", "", "Extensiones a excluir separadas por comas")

	flag.Parse()
	if *helpFlag {
		flag.Usage()
		return
	}
	if *excludedRoutesFlag != "" {
		excludedRoutes = strings.Split(*excludedRoutesFlag, ",")
	}
	if *fileExtensionsFlag != "" {
		fileExtensions = strings.Split(*fileExtensionsFlag, ",")
	}
	if *excludedFileExtensionsFlag != "" {
		excludedFileExtensions = strings.Split(*excludedFileExtensionsFlag, ",")
	}
	customFlags := hashes.CustomFlags{
		ShowHiddenFiles:        *showHiddenFiles,
		ExcludeRoutes:          excludedRoutes,
		FileExtensions:         fileExtensions,
		ExcludedFileExtensions: excludedFileExtensions,
	}

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
