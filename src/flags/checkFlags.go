package customFlags

import "strings"

type CustomFlags struct {
	ShowHiddenFiles        bool
	ExcludeRoutes          []string
	FileExtensions         []string
	ExcludedFileExtensions []string
}

func (cf *CustomFlags) Normalize() {
	cf.ExcludeRoutes = cleanSlice(cf.ExcludeRoutes)
	if len(cf.ExcludeRoutes) == 0 {
		cf.ExcludeRoutes = nil
	}

	cf.FileExtensions = cleanSlice(cf.FileExtensions)
	if len(cf.FileExtensions) == 0 {
		cf.FileExtensions = nil
	}

	cf.ExcludedFileExtensions = cleanSlice(cf.ExcludedFileExtensions)
	if len(cf.ExcludedFileExtensions) == 0 {
		cf.ExcludedFileExtensions = nil
	}
}

func cleanSlice(s []string) []string {
	var cleaned []string
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v != "" {
			cleaned = append(cleaned, v)
		}
	}
	return cleaned
}
