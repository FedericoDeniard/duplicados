package customFlags

import "strings"

type CustomFlags struct {
	ShowHiddenFiles        bool
	ExcludedRoutes         []string
	FileExtensions         []string
	ExcludedFileExtensions []string
	UseSHA256              bool
}

func (cf *CustomFlags) Normalize() {
	cf.ExcludedRoutes = cleanSlice(cf.ExcludedRoutes)
	if len(cf.ExcludedRoutes) == 0 {
		cf.ExcludedRoutes = nil
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
