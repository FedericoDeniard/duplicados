package hashes

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"sync"
)

type CustomFlags struct {
	ShowHiddenFiles bool
	ExcludeRoutes   []string
	FileExtensions  []string
}

type FileHash struct {
	Path string
	MD5  string
}

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

func HashFiles(root string, flags CustomFlags) map[string][]string {
	results := make(map[string][]string)
	resultChan := make(chan FileHash, 1000)
	mu := &sync.Mutex{}
	pool := NewWorkerPool(50)
	pool.Run()

	pool.AddTask(&FolderHashTask{route: root, pool: pool, resultChan: resultChan, flags: flags})

	go func() {
		pool.wg.Wait()
		close(resultChan)
	}()

	for hash := range resultChan {
		mu.Lock()
		results[hash.MD5] = append(results[hash.MD5], hash.Path)
		mu.Unlock()
	}
	return results
}

type Task interface {
	Process() FileHash
}

type WorkerPool struct {
	concurrency int
	tasksChan   chan Task
	wg          *sync.WaitGroup
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.wg.Add(1)
	go func() {
		wp.tasksChan <- task
	}()
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}
}

func (wp *WorkerPool) Close() {
	close(wp.tasksChan)
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		tasksChan:   make(chan Task, concurrency),
		concurrency: concurrency,
		wg:          &sync.WaitGroup{},
	}
}

type FileHashTask struct {
	file       *os.File
	resultChan chan FileHash
}

func (f *FileHashTask) Process() FileHash {
	md5, err := CalculateMD5(f.file)
	if err != nil {
		return FileHash{}
	}
	path, err := filepath.Abs(f.file.Name())
	if err != nil {
		path = f.file.Name()
	}
	hash := FileHash{Path: path, MD5: md5}
	f.resultChan <- hash
	return hash
}

type FolderHashTask struct {
	route      string
	pool       *WorkerPool
	resultChan chan FileHash
	flags      CustomFlags
}

func (f *FolderHashTask) Process() FileHash {
	files, err := os.ReadDir(f.route)
	if err != nil {
		return FileHash{}
	}

	for _, file := range files {
		path := filepath.Join(f.route, file.Name())
		if file.IsDir() {
			if !f.flags.ShowHiddenFiles && file.Name()[0] == '.' {
				continue
			}
			if slices.Contains(f.flags.ExcludeRoutes, file.Name()) || slices.Contains(f.flags.ExcludeRoutes, path) {
				continue
			}
			f.pool.AddTask(&FolderHashTask{route: path, pool: f.pool, resultChan: f.resultChan, flags: f.flags})
		} else {
			file, err := os.Open(path)
			if err != nil {
				continue
			}
			if file.Name()[0] == '.' {
				continue
			}
			fmt.Println(f.flags.FileExtensions, filepath.Ext(file.Name()))
			if len(f.flags.FileExtensions) > 0 && !slices.Contains(f.flags.FileExtensions, filepath.Ext(file.Name())) {
				continue
			}
			task := &FileHashTask{file: file, resultChan: f.resultChan}
			f.pool.AddTask(task)

		}
	}
	return FileHash{}
}
