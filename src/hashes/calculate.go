package hashes

import (
	"crypto/md5"
	"crypto/sha256"
	customFlags "duplicate-files/src/flags"
	"duplicate-files/src/types"
	"duplicate-files/src/workerpool"
	"fmt"
	"io"
	"os"
	"sync"
)

// Es más ligera
func CalculateMD5(file *os.File) (string, error) {
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Es más pesada pero más segura
func CalculateSHA256(file *os.File) (string, error) {
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// This function is unused
func GroupByHashes(files []types.FileHash) map[string][]string {
	duplicates := make(map[string][]string)
	for _, file := range files {
		duplicates[file.MD5] = append(duplicates[file.MD5], file.Path)
	}
	return duplicates
}

func HashFiles(root string, flags customFlags.CustomFlags) map[string][]string {
	results := make(map[string][]string)
	resultChan := make(chan types.FileHash, 1000)
	mu := &sync.Mutex{}
	pool := workerpool.NewWorkerPool(50)
	pool.Run()

	createTask := func(file *os.File, resultChan chan types.FileHash) *workerpool.FileHashTask {
		hashFunc := CalculateMD5
		if flags.UseSHA256 {
			hashFunc = CalculateSHA256
		}
		return workerpool.NewFileHashTask(file, resultChan, hashFunc)
	}

	folderTask := workerpool.NewFolderHashTask(root, pool, resultChan, flags, createTask)
	pool.AddTask(folderTask)

	go func() {
		pool.Wait()
		close(resultChan)
		pool.Close()
	}()

	for hash := range resultChan {
		mu.Lock()
		results[hash.MD5] = append(results[hash.MD5], hash.Path)
		mu.Unlock()
	}

	return results
}
