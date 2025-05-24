package workerpool

import (
	customFlags "duplicate-files/src/flags"
	"duplicate-files/src/types"
	"os"
	"path/filepath"
	"slices"
	"sync"
)

type Task interface {
	Process() types.FileHash
}

type FileHashTask struct {
	file        *os.File
	resultChan  chan types.FileHash
	hashFunc    func(*os.File) (string, error)
	getFilePath func(*os.File) (string, error)
}

func NewFileHashTask(file *os.File, resultChan chan types.FileHash, hashFunc func(*os.File) (string, error)) *FileHashTask {
	return &FileHashTask{
		file:        file,
		resultChan:  resultChan,
		hashFunc:    hashFunc,
		getFilePath: defaultGetFilePath,
	}
}

func (f *FileHashTask) Process() types.FileHash {
	md5, err := f.hashFunc(f.file)
	if err != nil {
		return types.FileHash{}
	}
	path, err := f.getFilePath(f.file)
	if err != nil {
		path = f.file.Name()
	}
	hash := types.FileHash{Path: path, MD5: md5}
	f.resultChan <- hash
	return hash
}

type FolderHashTask struct {
	Route      string
	Pool       *WorkerPool
	ResultChan chan types.FileHash
	Flags      customFlags.CustomFlags
	createTask func(*os.File, chan types.FileHash) *FileHashTask
}

func NewFolderHashTask(route string, pool *WorkerPool, resultChan chan types.FileHash, flags customFlags.CustomFlags, createTaskFunc func(*os.File, chan types.FileHash) *FileHashTask) *FolderHashTask {
	return &FolderHashTask{
		Route:      route,
		Pool:       pool,
		ResultChan: resultChan,
		Flags:      flags,
		createTask: createTaskFunc,
	}
}

func (f *FolderHashTask) Process() types.FileHash {
	files, err := os.ReadDir(f.Route)
	if err != nil {
		return types.FileHash{}
	}

	for _, file := range files {
		path := filepath.Join(f.Route, file.Name())
		if file.IsDir() {
			if !f.Flags.ShowHiddenFiles && file.Name()[0] == '.' {
				continue
			}
			if slices.Contains(f.Flags.ExcludeRoutes, file.Name()) || slices.Contains(f.Flags.ExcludeRoutes, path) {
				continue
			}
			f.Pool.AddTask(NewFolderHashTask(path, f.Pool, f.ResultChan, f.Flags, f.createTask))
		} else {
			fileIsNotInFileExtensions := len(f.Flags.FileExtensions) > 0 && !slices.Contains(f.Flags.FileExtensions, filepath.Ext(file.Name()))
			fileIsInExcludedFileExtensions := len(f.Flags.ExcludedFileExtensions) > 0 && slices.Contains(f.Flags.ExcludedFileExtensions, filepath.Ext(file.Name()))
			file, err := os.Open(path)
			if err != nil {
				continue
			}
			if file.Name()[0] == '.' {
				continue
			}
			if fileIsNotInFileExtensions || fileIsInExcludedFileExtensions {
				continue
			}
			task := f.createTask(file, f.ResultChan)
			f.Pool.AddTask(task)
		}
	}
	return types.FileHash{}
}

type WorkerPool struct {
	concurrency int
	tasksChan   chan Task
	wg          *sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		tasksChan:   make(chan Task, concurrency),
		concurrency: concurrency,
		wg:          &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}
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

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool) Close() {
	close(wp.tasksChan)
}

func defaultGetFilePath(file *os.File) (string, error) {
	return filepath.Abs(file.Name())
}
