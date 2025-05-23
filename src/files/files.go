package filesManager

import "os"

type DirClassified struct {
	Files  []string
	Routes []string
	Error  error
}

func ClassifyFilesOrRoutes(files []os.DirEntry) DirClassified {
	var filePaths, routesPaths []string
	for _, file := range files {
		if !file.IsDir() {
			filePaths = append(filePaths, file.Name())
		} else {
			routesPaths = append(routesPaths, file.Name())
		}
	}

	return DirClassified{Files: filePaths, Routes: routesPaths}
}
