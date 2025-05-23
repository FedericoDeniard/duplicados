package main

import (
	"duplicate-files/src/hashes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	basePath := filepath.Join("/home/federico/Documentos/Programación/")
	files, err := os.ReadDir(basePath)
	if err != nil {
		panic(err)
	}

	var filesHashes []hashes.FileHash
	var wg sync.WaitGroup
	var mu sync.Mutex

	hashes.HashFiles(files, basePath, &filesHashes, &wg, &mu)
	wg.Wait()
	fmt.Println(len(filesHashes))

	// classified := filesManager.ClassifyFilesOrRoutes(files)
	// if classified.Error != nil {
	// 	panic(classified.Error)
	// }

	// var hasedFiles []hashes.FileHash
	// for _, file := range classified.Files {
	// 	md5, err := hashes.CalculateMD5("./pruebas/" + file)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	hasedFiles = append(hasedFiles, hashes.FileHash{Path: "./pruebas/" + file, MD5: md5})
	// }

	// // Buscar duplicados
	// groupsByHashes := hashes.GroupByHashes(hasedFiles)
	// fmt.Println("Duplicados")
	// for _, paths := range groupsByHashes {
	// 	if len(paths) > 1 {
	// 		fmt.Println(paths)
	// 	}
	// }
}
