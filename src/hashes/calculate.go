package hashes

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type FileHash struct {
	Path string
	MD5  string
}

// Es más ligera
func CalculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Es más pesada pero más segura
func CalculateSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func GroupByHashes(files []FileHash) map[string][]string {
	duplicates := make(map[string][]string)
	for _, file := range files {
		if _, ok := duplicates[file.MD5]; ok {
			duplicates[file.MD5] = append(duplicates[file.MD5], file.Path)
		} else {
			duplicates[file.MD5] = []string{file.Path}
		}
	}
	return duplicates
}

func HashFiles(files []os.DirEntry, basePath string, filesHashes *[]FileHash, wg *sync.WaitGroup, mu *sync.Mutex) {
	for _, file := range files {
		route := filepath.Join(basePath, file.Name())
		if !file.IsDir() {
			wg.Add(1)
			go func(route string) {
				defer wg.Done()
				md5, err := CalculateMD5(route)
				if err != nil {
					return
				}
				mu.Lock()
				*filesHashes = append(*filesHashes, FileHash{Path: route, MD5: md5})
				mu.Unlock()
			}(route)
		} else {
			wg.Add(1)
			go func(route string) {
				defer wg.Done()
				subFiles, err := os.ReadDir(route)
				if err != nil {
					return
				}
				HashFiles(subFiles, route, filesHashes, wg, mu)
			}(route)
		}
		// fmt.Println(route)
	}
}
